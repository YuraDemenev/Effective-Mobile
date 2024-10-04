package handlers

import "effective_mobile/elements"

//Структуры для фунции Info
//Структура для возвращения ответа
type returnSongInfoStruct struct {
	RequestData string
	Text        string
	Link        string
}

//Структура на получение json из запроса
type getSongInfoStruct struct {
	Group string `json:"group"`
	Name  string `json:"song"`
}

//Структуры для фунции createNewSong
// Структура для получения json из post запроса на создание песни
type CreateNewSongStruct struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	Link        string `json:"link"`
	Text        string `json:"text"`
	RequestData string `json:"release_date"`
}

//Структуры для функции getAllSongs
// Структура для получения всех песен, для обработки json
type getAllSongsStruct struct {
	Direction        string `json:"direction"`
	Page             int    `json:"page"`
	Filter           string `json:"filter"`
	Field            string `json:"field"`
	CountSongsOnPage int    `json:"count_songs_on_page"`
}

//Структура для возвращения ответа
type returnAllSongsStruct struct {
	Songs      []elements.SongStruct
	CountPages int
}

//Структуры для функции deleteSong
//Структура для возвращения ответа
type returnDeleteSongStruct struct {
	Result string
}

//Структуры для фунции changeSong
// Структура для получения json из post запроса на создание песни
type getChangeSongStruct struct {
	Id          int    `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	Link        string `json:"link"`
	Text        string `json:"text"`
	RequestData string `json:"release_date"`
}

//Структура для функции getSongByVerses
//Структура для возвращения ответа
type returnGetSongByVersesStruct struct {
	Song       elements.SongStruct
	Verses     []string
	CountPages int
}
