package domain

import "time"

type Music struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	File      []byte    `json:"file"`
	CreatedAt time.Time `json:"created_at"`
}

func (Music) TableName() string {
	return "music"
}

type MusicUsecase interface {
	GetAllMusic() ([]*Music, error)
	AddMusic(music *Music) error
	UpdateMusic(music *Music) error
	DeleteMusic(id string) error
	GetMusicByID(id uint64) (*Music,error)
}

type MusicRepository interface {
	GetAll() ([]*Music, error)
	Add(music *Music) error
	Update(music *Music) error
	Delete(id string) error
	GetByID(id uint64) (*Music, error)
}
