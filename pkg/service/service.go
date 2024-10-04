package service

import (
	"effective_mobile/elements"
	"effective_mobile/pkg/repository"
)

type Service struct {
	Songs
}

type Songs interface {
	SaveNewSong(group, song, link, text, releaseData string) error
	GetAllSongs(direction, field, filter string, page, countSongsOnPage int) (songs []elements.SongStruct,
		countPages, errorCode int, err error)
	GetSongInfo(group, name string) (releaseDate, text, link string, id int, err error)
	DeleteSong(group, song string) error
	ChangeSong(id int, group, song, link, text, releaseData string) error
	GetSongByVerses(countVersesOnPages, page int, group, song, direction string) (verses []string,
		countPages, songId, errorCode int, err error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Songs: NewSongsService(repos.Songs)}
}
