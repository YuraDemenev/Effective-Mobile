package elements

// Структура песни для отправки ответа в get_all_songs
type SongStruct struct {
	Group  string `json:"group"`
	Song   string `json:"song"`
	SongId int    `json:"song_id"`
}

// Структура для получения песни по куплетам с пагинацией
type GetSongByVerseStruct struct {
	Group              string `json:"group"`
	Song               string `json:"song"`
	CountVersesOnPages int    `json:"count_verses_on_pages"`
	Page               int    `json:"page"`
	Direction          string `json:"direction"`
}
