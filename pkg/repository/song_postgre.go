package repository

import (
	"effective_mobile/elements"
	"errors"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
)

// В этом файле функции для взаимодейтсвиями с tokens в БД
type SongPostgre struct {
	db *sqlx.DB
}

func NewSongPostgre(db *sqlx.DB) *SongPostgre {
	return &SongPostgre{db: db}
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------------------------

// Функция на сохранение новой песни
func (s *SongPostgre) SaveSong(group, song, link string, verses []string, releaseDate string) error {
	var groupId int
	var songId int

	//Проверяем есть ли такая группа
	query_ := fmt.Sprintf(`
	SELECT id
	FROM %s
	WHERE EXISTS (
		SELECT 1
		FROM %s
		WHERE %s.name = '%s'
	);`, groupsTable, groupsTable, groupsTable, group)
	err := s.db.QueryRow(query_).Scan(&groupId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}

	//Если нет то создаём новую запись
	if groupId == 0 {
		//Запрос на сохранение данных в groups
		query_ = fmt.Sprintf(`
		INSERT INTO %s (name)
		VALUES ($1)
		RETURNING id;`, groupsTable)

		err := s.db.QueryRow(query_, group).Scan(&groupId)
		if err != nil {
			return err
		}
	}

	//Запрос на сохранение данных в songs
	query_ = fmt.Sprintf(`
	INSERT INTO %s (group_id,name,link,release_date) 
	VALUES ($1,$2,$3,$4) 
	RETURNING id;
	`, songsTable)

	err = s.db.QueryRow(query_, groupId, song, link, releaseDate).Scan(&songId)
	if err != nil {
		return err
	}

	//Запрос на сохранение данных в verses
	for i, vers := range verses {
		query_ = fmt.Sprintf(`
		INSERT INTO %s(song_id,verse_number,text)
		VALUES ($1,$2,$3)
		`, versesTable)

		err = s.db.QueryRow(query_, songId, i+1, vers).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------------------------

// Функция на получения всех песен с заданными параметрами
func (s *SongPostgre) GetAllSongs(direction, field, filter string, page, countSongsOnPage int) (songs []elements.SongStruct,
	countPages, errorCode int, err error) {
	var countSongs int

	//Константы для проверки field и filter
	const filterFieldName = "name"
	const filterFieldGroup = "group"
	const filterASC = "asc"
	const filterDESC = "desc"
	//Для проверки direction
	const directionNext = "next"
	const directionPrevious = "previous"

	checkFilter := false
	lastId := 0

	//Запрос на получение кол-во песен
	query_ := fmt.Sprintf(`
		SELECT COUNT (*)
		FROM %s;
	`, songsTable)
	err = s.db.QueryRow(query_).Scan(&countSongs)
	if err != nil {
		return nil, 0, http.StatusInternalServerError, err
	}

	//Проверяем наличие поля по которому нужно отвильтровать
	if field != "" || filter != "" {
		//Проверяем корректность предоставленных данных фильтров
		if field != filterFieldGroup && field != filterFieldName {
			errorStr := fmt.Sprintf("wrong data in filter fields. fields is not %s/%s", filterFieldGroup, filterFieldName)
			return nil, 0, http.StatusBadRequest, errors.New(errorStr)
		}
		//Проверяем корректность данных по тому как фильтровать
		if filter != filterASC && filter != filterDESC {
			errorStr := fmt.Sprintf("wrong data in filter. filter is not %s/%s", filterASC, filterDESC)
			return nil, 0, http.StatusBadRequest, errors.New(errorStr)
		}
		checkFilter = true
	}

	if direction == "" {
		//Проверяем верно ли введен номер страницы
		if page <= 0 || page > int(math.Ceil(float64(countSongs)/float64(countSongsOnPage))) {
			return nil, 0, http.StatusBadRequest, errors.New("page number not correct")
		}
		//Получаем id последней песни на предыдущей страницы
		lastId = (page - 1) * countSongsOnPage
		//Если введен direction
	} else {
		//Проверка верны ли данные
		if direction != directionNext && direction != directionPrevious {
			return nil, 0, http.StatusBadRequest, errors.New("not correct direction")
		}
		//Если мы запрашиваем предыдущую страницу
		if direction == directionPrevious {
			if page-2 > 0 {
				lastId = (page - 2) * countSongsOnPage
			} else if page-2 == 0 {
				lastId = countSongsOnPage
			} else {
				return nil, 0, http.StatusBadRequest, errors.New("can`t use previous direction")
			}
			//Если следующую нам нужно вычислить id последней песни на нашей странце
		} else {
			lastId = (page) * countSongsOnPage
			//Проверяем что last id возможен
			if lastId >= countSongs {
				return nil, 0, http.StatusBadRequest, errors.New("can`t use next direction")
			}
		}
	}

	//Запрос на полуение песен (получаем столько сколько указанно в запросе)
	query_ = fmt.Sprintf(`
		SELECT groups.name as group_name,songs.name as song_name, songs.id as song_id
		FROM %s
		JOIN %s on songs.group_id = groups.id
		WHERE %s.id > %d
		ORDER BY %s.id ASC
		LIMIT %d
	`, songsTable, groupsTable, songsTable, lastId, songsTable, countSongsOnPage)

	//Делаем запрос на получение данных из DB
	rows, err := s.db.Query(query_)
	if err != nil {
		return nil, 0, http.StatusInternalServerError, err
	}
	defer rows.Close()

	//Проходим по результатом и добавляем всё в финальный массив
	result := make([]elements.SongStruct, 0, countSongsOnPage)
	for rows.Next() {
		var song elements.SongStruct
		if err := rows.Scan(&song.Group, &song.Song, &song.SongId); err != nil {
			return nil, 0, http.StatusInternalServerError, err
		}
		result = append(result, song)
	}

	//Если нужно применить фильтр, применяем к результатам
	if checkFilter {
		//Отфильтровывем результат
		sort.Slice(result, func(i, j int) bool {
			//Если группы
			if field == filterFieldGroup {
				if filter == filterASC {
					return result[i].Group < result[j].Group
				} else {
					return result[i].Group > result[j].Group
				}
			} else {
				if filter == filterASC {
					return result[i].Song < result[j].Song
				} else {
					return result[i].Song > result[j].Song
				}
			}
		})
	}

	return result, int(math.Ceil(float64(countSongs) / float64(countSongsOnPage))), 0, nil
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------------------------

// Получение информации о песни
func (s *SongPostgre) GetSongInfo(group, name string) (releaseDate, text, link string, id int, err error) {
	var songId int

	//Запрос на информации о песни
	query_ := fmt.Sprintf(`
		SELECT songs.release_date,songs.id, songs.link
		FROM %s
		JOIN %s on songs.group_id = groups.id
		WHERE %s.name = '%s' AND %s.name = '%s';
	`, songsTable, groupsTable, songsTable, name, groupsTable, group)

	//Делаем запрос
	err = s.db.QueryRow(query_).Scan(&releaseDate, &songId, &link)
	if err != nil {
		return "", "", "", 0, err
	}

	//Запрос на получени текста из таблицы verses
	query_ = fmt.Sprintf(`
		SELECT text
		FROM %s
		WHERE song_id = %d
	`, versesTable, songId)

	rows, err := s.db.Query(query_)
	if err != nil {
		return "", "", "", 0, err
	}
	defer rows.Close()

	var stringBuilder strings.Builder
	//Проходим по результатом и добавляем всё в финальную строку
	for rows.Next() {
		var str string
		if err := rows.Scan(&str); err != nil {
			return "", "", "", 0, err
		}
		stringBuilder.WriteString(str)
	}

	return releaseDate, stringBuilder.String(), link, songId, nil
}

// --------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------------------------------------------------------
// Функция на удаление песни
func (s *SongPostgre) DeleteSong(group, song string) error {
	var groupId int
	var count int
	//Удаление песни
	query_ := fmt.Sprintf(`
		DELETE FROM %s
		WHERE %s.name = '%s'
		RETURNING group_id;
	`, songsTable, songsTable, song)
	//Делаем запрос
	err := s.db.QueryRow(query_).Scan(&groupId)
	if err != nil {
		return err
	}

	//Получаем кол-во песен у которых такой же group_id
	query_ = fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s
		WHERE group_id = %d
	`, songsTable, groupId)
	err = s.db.QueryRow(query_).Scan(&count)
	if err != nil {
		return err
	}

	//Если кол-во 0 удаляем из таблицы groups группу
	if count == 0 {
		query_ = fmt.Sprintf(`
			DELETE FROM %s
			WHERE id = %d
			RETURNING id;
		`, groupsTable, groupId)

		err = s.db.QueryRow(query_).Scan(&groupId)
		if err != nil {
			return err
		}
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------------------------------------------------------
// Функция для изменения песни
func (s *SongPostgre) ChangeSong(id int, group, song, link, releaseDate string, verses []string) error {
	var songId int
	var group_id int

	//Если мы вносим изменения в текст
	if len(verses) != 0 {
		//Удаляем из таблицы verses записи текущей песней
		query_ := fmt.Sprintf(`
		DELETE FROM %s
		WHERE song_id = %d
		RETURNING id
		`, versesTable, id)
		err := s.db.QueryRow(query_).Scan(&songId)
		if err != nil {
			return err
		}

		//Запрос на сохранение данных в verses
		for i, vers := range verses {
			query_ = fmt.Sprintf(`
			INSERT INTO %s(song_id,verse_number,text)
			VALUES ($1,$2,$3)
			`, versesTable)

			err = s.db.QueryRow(query_, id, i+1, vers).Err()
			if err != nil {
				return err
			}
		}
	}

	//Если мы вносим изменения в группу
	if group != "" {
		//Проверяем есть ли группа с тем именем на которое мы хотим поменять в таблице group и возвращаем id
		query_ := fmt.Sprintf(`
		SELECT id
		FROM %s
		WHERE name = '%s'
		`, groupsTable, group)
		err := s.db.QueryRow(query_).Scan(&group_id)
		if err != nil {
			return err
		}
		//Для того чтобы после обновления проверить сколько песен с этим group_id, если 0 удалить группу из DB
		checkGroupId := -1

		//Если group_id = 0 значит такой группы нет, надо её создать и обновить в таблице songs group_id
		if group_id == 0 {
			//Запрос на сохранение данных в groups
			var newGroupId int
			query_ = fmt.Sprintf(`
			INSERT INTO %s (name)
			VALUES ($1)
			RETURNING id;`, groupsTable)

			err := s.db.QueryRow(query_, group).Scan(&newGroupId)
			if err != nil {
				return err
			}
			//Обновляем group_id в songs
			query_ = fmt.Sprintf(`	
			UPDATE %s
			SET group_id = %d
			WHERE id = %d
			`, songsTable, newGroupId, id)
			err = s.db.QueryRow(query_).Err()
			if err != nil {
				return err
			}
			checkGroupId = group_id

		} else {
			//Получаем id предыдущей группы
			var previousGroupId int
			query_ = fmt.Sprintf(`
			SELECT group_id
			FROM %s
			WHERE id = %d;
			`, songsTable, id)
			err = s.db.QueryRow(query_).Scan(&previousGroupId)
			if err != nil {
				return err
			}

			//Заменяем старую id группы на новую
			query_ = fmt.Sprintf(`	
			UPDATE %s
			SET group_id = %d
			WHERE id = %d
			`, songsTable, group_id, id)
			err = s.db.QueryRow(query_).Err()
			if err != nil {
				return err
			}
			checkGroupId = previousGroupId
		}
		//Получаем кол-во песен с у которых group_id = checkGroupId (если кол-во 0 группу нужно удалить)
		var countSongs int
		query_ = fmt.Sprintf(`
		SELECT COUNT (*)
		FROM %s
		WHERE %s.group_id = %d
		`, songsTable, songsTable, checkGroupId)
		err = s.db.QueryRow(query_).Scan(&countSongs)
		if err != nil {
			return err
		}

		//Если кол-во песен равно 0, удаляем группу
		if countSongs == 0 {
			query_ = fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = %d
		`, groupsTable, checkGroupId)
			err = s.db.QueryRow(query_).Err()
			if err != nil {
				return err
			}
		}
	}

	//Для проверки что нужно изменить link/name/release_date
	checkHasChanges := false
	//Заменяем данные в таблице songs
	query_ := fmt.Sprintf(`
		UPDATE %s
		SET 
	`, songsTable)
	//Меняем ссылку
	if link != "" {
		query_ += fmt.Sprintf("link = '%s',", link)
		checkHasChanges = true
	}
	//Меняем название песни
	if song != "" {
		query_ += fmt.Sprintf("name = '%s',", song)
		checkHasChanges = true
	}
	//Меняем дату релиза
	if releaseDate != "" {
		query_ += fmt.Sprintf("release_date = '%s',", releaseDate)
		checkHasChanges = true
	}
	if checkHasChanges {
		//Чтобы удалить в конце ","
		query_ = query_[0 : len(query_)-1]
		query_ += fmt.Sprintf(" WHERE id = %d", id)
		err := s.db.QueryRow(query_).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------------------------------------------------------
// Функция для получения песен по куплетам
func (s *SongPostgre) GetSongByVerses(countVersesOnPages, page int, group, song, direction string) (verses []string, countPages,
	songId, errorCode int, err error) {
	var countVerse int
	//Для проверки direction
	const directionNext = "next"
	const directionPrevious = "previous"

	lastVerseId := 0
	songId = 0
	//Запрос на получение id песни
	query_ := fmt.Sprintf(`
	SELECT songs.id
		FROM %s
		JOIN %s on songs.group_id = groups.id
		WHERE %s.name = '%s' AND %s.name ='%s';
	`, songsTable, groupsTable, songsTable, song, groupsTable, group)
	err = s.db.QueryRow(query_).Scan(&songId)
	if err != nil {
		return
	}

	//Запрос на получение кол-во куплетов
	query_ = fmt.Sprintf(`
		SELECT COUNT (*)
		FROM %s
		WHERE song_id = %d;
	`, versesTable, songId)
	err = s.db.QueryRow(query_).Scan(&countVerse)
	if err != nil {
		return
	}

	if direction == "" {
		//Проверяем верно ли введен номер страницы
		if page <= 0 || page > int(math.Ceil(float64(countVerse)/float64(countVersesOnPages))) {
			return nil, 0, 0, http.StatusBadRequest, errors.New("page number not correct")
		}
		//Получаем id последней песни на предыдущей страницы
		lastVerseId = (page - 1) * countVersesOnPages
		//Если введен direction
	} else {
		//Проверка верны ли данные
		if direction != directionNext && direction != directionPrevious {
			return nil, 0, 0, http.StatusBadRequest, errors.New("not correct direction")
		}
		//Если мы запрашиваем предыдущую страницу
		if direction == directionPrevious {
			if page-2 > 0 {
				lastVerseId = (page - 2) * countVersesOnPages
			} else if page-2 == 0 {
				lastVerseId = countVersesOnPages
			} else {
				return nil, 0, 0, http.StatusBadRequest, errors.New("can`t use previous direction")
			}
			//Если следующую нам нужно вычислить id последней песни на нашей странце
		} else {
			lastVerseId = (page) * countVersesOnPages
			//Проверяем что last id возможен
			if lastVerseId >= countVerse {
				return nil, 0, 0, http.StatusBadRequest, errors.New("can`t use next direction")
			}
		}
	}

	//Запрос на получение куплетов (получаем столько сколько указанно в запросе)
	query_ = fmt.Sprintf(`
		SELECT text
		FROM %s
		WHERE %s.id > %d
		LIMIT %d
	`, versesTable, versesTable, lastVerseId, countVersesOnPages)

	//Делаем запрос на получение данных из DB
	rows, err := s.db.Query(query_)
	if err != nil {
		return nil, 0, 0, http.StatusInternalServerError, err
	}
	defer rows.Close()

	//Проходим по результатом и добавляем всё в финальный массив
	verses = make([]string, 0, countVersesOnPages)
	for rows.Next() {
		var str string
		if err := rows.Scan(&str); err != nil {
			return nil, 0, 0, http.StatusInternalServerError, err
		}
		verses = append(verses, str)
	}

	return verses, int(math.Ceil(float64(countVerse) / float64(countVersesOnPages))), songId, 0, nil
}
