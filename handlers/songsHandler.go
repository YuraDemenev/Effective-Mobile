package handlers

import (
	"bytes"
	"effective_mobile/elements"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SetErrorStuct struct {
	Error string `json:"error"`
	Info  string `json:"info"`
}

// Словарь чтобы просто и быстро получать строку ответа по коду ошибки
var DictionaryErrors = map[int]string{
	400: "Неккорентные данные",
	500: "Внутренняя ошибка сервера",
}

// Функция для создания новой песни
// @Summary Create new song
// @Tags Songs
// @Desctiption create new song with params
// @Accept json
// @Produce json
// @Param input body CreateNewSongStruct true "Функция создаёт новую песню из входящих данных"
// @Success 200 {object} returnSongInfoStruct
// @Failure 400,500 {object} SetErrorStuct
// @Router /create_new_song [post]
func (h *Handlers) createNewSong(c *gin.Context) {
	//Структура для получения данных
	var requestData CreateNewSongStruct

	//Функция для декодирования запроса
	if !DecodeRequest(&requestData, c) {
		return
	}

	//Провереяем наличие пустых полей
	if requestData.Group == "" || requestData.Song == "" || requestData.RequestData == "" ||
		requestData.Link == "" || requestData.Text == "" {
		WriteError(c, "Request has empty field", http.StatusBadRequest)
		return
	}

	//Сохраняем данные
	err := h.service.Songs.SaveNewSong(requestData.Group, requestData.Song, requestData.Link, requestData.Text, requestData.RequestData)
	if err != nil {
		WriteError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	//Подготавливаем get запрос
	// Создаем JSON с информацией о группе и песне
	requestBody, err := json.Marshal(map[string]string{
		"group": requestData.Group,
		"song":  requestData.Song,
	})
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	//В случае запуска теста c.Request.Host возвращает "", поэтому делаем проверку
	host := c.Request.Host
	if host == "" {
		host = "localhost:8080"
	}
	url := "http://" + host + "/info"
	//Создаём запрос
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
	if err != nil {
		WriteError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	//Пишем хедер
	req.Header.Set("Content-Type", "application/json")
	//Делаем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		WriteError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//Получаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		WriteError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Write(body)
}

// Функция на получение всех песен
// @Summary Get all songs
// @Tags Songs
// @Desctiption get all songs with pagination
// @Accept json
// @Produce json
// @Param input body getAllSongsStruct true "Фунция возвращает все песни с пагинацией"
// @Success 200 {object} returnAllSongsStruct
// @Failure 400,500 {object} SetErrorStuct
// @Router /get_all_songs [get]
func (h *Handlers) getAllSongs(c *gin.Context) {
	//Структура для получения данных
	var requestData getAllSongsStruct

	//Функция для декодирования запроса
	if !DecodeRequest(&requestData, c) {
		return
	}
	//Проверяем данные
	if requestData.CountSongsOnPage <= 0 {
		WriteError(c, "count songs on page is not valid", http.StatusBadRequest)
		return
	}

	//Получаем все песни на странице
	songs, countPages, statusError, err := h.service.Songs.GetAllSongs(requestData.Direction, requestData.Field, requestData.Filter,
		requestData.Page, requestData.CountSongsOnPage)
	if err != nil {
		//Возвращаем ошибку
		WriteError(c, err.Error(), statusError)
		return
	}

	returnValue := returnAllSongsStruct{Songs: songs, CountPages: countPages}
	//Конвертруем в json  заранее заготовленную структуры, для возвращения ответа
	result, err := json.Marshal(returnValue)
	if err != nil {
		str := fmt.Sprintf("Can`t convert returnValueStruct to json. Error: %s", err.Error())
		//Возвращаем ошибку
		WriteError(c, str, http.StatusInternalServerError)
		return
	}

	c.Writer.Write(result)
}

// Функция на получение информации о конкретной песни
// @Summary Info
// @Tags Songs
// @Desctiption get info about song
// @Accept json
// @Produce json
// @Param input body getAllSongsStruct true "Фунция возвращает информацию о песни"
// @Success 200 {object} returnSongInfoStruct
// @Failure 400,500 {object} SetErrorStuct
// @Router /info [get]
func (h *Handlers) info(c *gin.Context) {
	var requestData getSongInfoStruct
	//Функция для декодирования запроса
	if !DecodeRequest(&requestData, c) {
		return
	}

	if requestData.Group == "" || requestData.Name == "" {
		WriteError(c, "not correct data. group or name is empty", http.StatusBadRequest)
		return
	}

	releaseDate, text, link, _, err := h.service.GetSongInfo(requestData.Group, requestData.Name)
	if err != nil {
		WriteError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	returnValue := returnSongInfoStruct{RequestData: releaseDate, Text: text, Link: link}
	//Конвертруем в json  заранее заготовленную структуры, для возвращения ответа
	result, err := json.Marshal(returnValue)
	if err != nil {
		str := fmt.Sprintf("Can`t convert returnValueStruct to json. Error: %s", err.Error())
		//Возвращаем ошибку
		WriteError(c, str, http.StatusInternalServerError)
		return
	}

	c.Writer.Write(result)
}

// Функция для удаления песни
// @Summary Delete song
// @Tags Songs
// @Desctiption delete song
// @Accept json
// @Produce json
// @Param input body getSongInfoStruct true "Удаление песни"
// @Success 200 {object} returnDeleteSongStruct
// @Failure 400,500 {object} SetErrorStuct
// @Router /delete_song [delete]
// Функция на получение информации о конкретной песни
func (h *Handlers) deleteSong(c *gin.Context) {
	var requestData getSongInfoStruct
	//Функция для декодирования запроса
	if !DecodeRequest(&requestData, c) {
		return
	}
	if requestData.Group == "" || requestData.Name == "" {
		WriteError(c, "group or name is empty", http.StatusBadRequest)
	}

	err := h.service.Songs.DeleteSong(requestData.Group, requestData.Name)
	if err != nil {
		WriteError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	returnValue := returnDeleteSongStruct{Result: "successed deleted"}
	//Конвертруем в json  заранее заготовленную структуры, для возвращения ответа
	result, err := json.Marshal(returnValue)
	if err != nil {
		str := fmt.Sprintf("Can`t convert returnValueStruct to json. Error: %s", err.Error())
		//Возвращаем ошибку
		WriteError(c, str, http.StatusInternalServerError)
		return
	}
	c.Writer.Write(result)
}

// Меняем песню
// @Summary Change song
// @Tags Songs
// @Desctiption Change song
// @Accept json
// @Produce json
// @Param input body getChangeSongStruct true "Меняем данные песни"
// @Success 200 {object} returnDeleteSongStruct
// @Failure 400,500 {object} SetErrorStuct
// @Router /change_song [patch]
// Функция на получение информации о конкретной песни
func (h *Handlers) changeSong(c *gin.Context) {
	var requestData getChangeSongStruct
	//Декодируеvм полученный json
	if !DecodeRequest(&requestData, c) {
		return
	}
	if (requestData.Group == "" && requestData.Link == "" && requestData.RequestData == "" && requestData.Text == "") ||
		requestData.Id <= 0 {
		WriteError(c, "Not correct data. All empty or id <=0", http.StatusBadRequest)
		return
	}
	err := h.service.Songs.ChangeSong(requestData.Id, requestData.Group, requestData.Song, requestData.Link,
		requestData.Text, requestData.RequestData)
	if err != nil {
		WriteError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	returnValue := returnDeleteSongStruct{Result: "successed changed"}
	//Конвертруем в json  заранее заготовленную структуры, для возвращения ответа
	result, err := json.Marshal(returnValue)
	if err != nil {
		str := fmt.Sprintf("Can`t convert returnValueStruct to json. Error: %s", err.Error())
		//Возвращаем ошибку
		WriteError(c, str, http.StatusInternalServerError)
		return
	}
	c.Writer.Write(result)
}

// Получаем песню по слогам с пагинацией
// @Summary get song by verses
// @Tags Songs
// @Desctiption get song with verses
// @Accept json
// @Produce json
// @Param input body elements.GetSongByVerseStruct true "Получаем песню с куплетами, с пагинацией по куплетам"
// @Success 200 {object} returnGetSongByVersesStruct
// @Failure 400,500 {object} SetErrorStuct
// @Router /get_song_by_verses [get]
func (h *Handlers) getSongByVerses(c *gin.Context) {
	var requestData elements.GetSongByVerseStruct
	//Декодируеvм полученный json
	if !DecodeRequest(&requestData, c) {
		return
	}

	if requestData.Group == "" || requestData.Song == "" {
		WriteError(c, "group/song is empty", http.StatusBadRequest)
		return
	}
	if requestData.CountVersesOnPages <= 0 {
		WriteError(c, "count verses on page is not valid", http.StatusBadRequest)
		return
	}
	verses, countPages, songId, errorCode, err := h.service.Songs.GetSongByVerses(requestData.CountVersesOnPages,
		requestData.Page, requestData.Group, requestData.Song, requestData.Direction)
	if err != nil {
		WriteError(c, err.Error(), errorCode)
		return
	}

	song := elements.SongStruct{Song: requestData.Song, Group: requestData.Group, SongId: songId}
	returnValue := returnGetSongByVersesStruct{Song: song, Verses: verses, CountPages: countPages}
	//Конвертруем в json  заранее заготовленную структуры, для возвращения ответа
	result, err := json.Marshal(returnValue)
	if err != nil {
		str := fmt.Sprintf("Can`t convert returnValueStruct to json. Error: %s", err.Error())
		//Возвращаем ошибку
		WriteError(c, str, http.StatusInternalServerError)
		return
	}
	c.Writer.Write(result)
}

// Функция для декодирования запроса
func DecodeRequest(requestData interface{}, c *gin.Context) bool {
	//Создаем декодер
	decoder := json.NewDecoder(c.Request.Body)
	//Запрещаем иметь в запросе поля которых нет в структуре
	decoder.DisallowUnknownFields()

	//Декодируем
	err := decoder.Decode(&requestData)
	if err != nil {
		str := fmt.Sprintf("Ошибка при декодирования request body. Ошибка: %s", err.Error())
		//Возвращаем ошибку
		WriteError(c, str, http.StatusBadRequest)
		return false
	}
	return true
}

// Функция для возвращения ошибки
func WriteError(c *gin.Context, str string, errorNum int) {
	logrus.Error(str)
	//Возвращаем ошибку в postman

	//Создаём структуру для возвращения ошибки
	var setErrorStuct SetErrorStuct
	setErrorStuct.Error = DictionaryErrors[errorNum]
	setErrorStuct.Info = str

	//Конвертируем в json
	result, err := json.Marshal(setErrorStuct)
	if err != nil {
		str := fmt.Sprintf("Can`t convert setErrorStuct to json. Error: %s", err.Error())
		logrus.Error(str)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Возвращаем ошибку
	c.Writer.WriteHeader(errorNum)
	c.Writer.Write(result)
}
