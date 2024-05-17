package domain

import "time"

type Music struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	File      []byte    `json:"file"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type MusicUsecase interface {
	GetAllMusic() ([]*Music, error)
	AddMusic(music *Music) error
	UpdateMusic(music *Music) error
	DeleteMusic(id string) error
	GetMusicById(id string) ([]*Music, error)
}

type MusicRepository interface {
	GetAll() ([]*Music, error)
	Add(music *Music) error
	Update(music *Music) error
	Delete(id string) error
	GetMusicById(id string) ([]*Music, error)
}
