package repository

import (
	"effective_mobile/elements"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Songs
}

type Songs interface {
	SaveSong(group, song, link string, verses []string, releaseData string) error
	GetAllSongs(direction, field, filter string, page, countSongsOnPage int) (songs []elements.SongStruct,
		countPages, errorCode int, err error)
	GetSongInfo(group, name string) (releaseDate, text, link string, id int, err error)
	DeleteSong(group, song string) error
	ChangeSong(id int, group, song, link, releaseData string, verses []string) error
	GetSongByVerses(countVersesOnPages, page int, group, song, direction string) (verses []string,
		countPages, songId, errorCode int, err error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Songs: NewSongPostgre(db),
	}
}
