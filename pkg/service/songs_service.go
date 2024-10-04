package service

import (
	"effective_mobile/elements"
	"effective_mobile/pkg/repository"
	"strings"
	"time"
)

type SongsService struct {
	repoSongs repository.Songs
}

func NewSongsService(repoSongs repository.Songs) *SongsService {
	return &SongsService{repoSongs: repoSongs}
}

// Generate Token
func (s *SongsService) SaveNewSong(group, song, link, text, releaseDate string) error {
	//Проверяем формат releaseDate
	layout := "02.01.2006"
	_, err := time.Parse(layout, releaseDate)
	if err != nil {
		return err
	}

	//Разделяем текст на абзацы
	verses := strings.Split(text, "/n")
	//Возварщаем результат
	return s.repoSongs.SaveSong(group, song, link, verses, releaseDate)
}

// Функция для получения всех песен
func (s *SongsService) GetAllSongs(direction, field, filter string, page, countSongsOnPage int) (songs []elements.SongStruct,
	countPages, errorCode int, err error) {
	return s.repoSongs.GetAllSongs(direction, field, filter, page, countSongsOnPage)
}

// Функция для получения информации об песни
func (s *SongsService) GetSongInfo(group, name string) (releaseDate, text, link string, id int, err error) {
	return s.repoSongs.GetSongInfo(group, name)
}

// Функция на удаление записи
func (s *SongsService) DeleteSong(group, song string) error {
	return s.repoSongs.DeleteSong(group, song)
}

// Функция для изменения песни
func (s *SongsService) ChangeSong(id int, group, song, link, text, releaseDate string) error {
	//Если надо менять releaseDate проверяем дату
	if releaseDate != "" {
		//Проверяем формат releaseDate
		layout := "02.01.2006"
		_, err := time.Parse(layout, releaseDate)
		if err != nil {
			return err
		}
	}
	verses := []string{}
	//Если text не пустой
	if text != "" {
		//Разделяем текст на абзацы
		verses = strings.Split(text, "/n")
	}
	//Возварщаем результат
	return s.repoSongs.ChangeSong(id, group, song, link, releaseDate, verses)
}

func (s *SongsService) GetSongByVerses(countVersesOnPages, page int, group, song, direction string) (verses []string,
	countPages, songId, errorCode int, err error) {
	return s.repoSongs.GetSongByVerses(countVersesOnPages, page, group, song, direction)
}
