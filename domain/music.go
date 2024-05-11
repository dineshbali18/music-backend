package domain

import "time"

type Music struct {
	ID        uint64    `json:"id"`
	SongName  string    `json:"song_name"`
	SongFile  []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type MusicUsecase interface {
	GetAllMusic() ([]*Music, error)
	AddMusic(music *Music) error
	UpdateMusic(music *Music) error
	DeleteMusic(id string) error
}

type MusicRepository interface {
	GetAll() ([]*Music, error)
	Add(music *Music) error
	Update(music *Music) error
	Delete(id string) error
}
