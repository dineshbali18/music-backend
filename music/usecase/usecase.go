package usecase

import (
	"fmt"
	"musicApp/domain"
)

type musicUsecase struct {
	musicRepository domain.MusicRepository
}

func NewMusicUsecase(repo domain.MusicRepository) domain.MusicUsecase {
	return &musicUsecase{musicRepository: repo}
}

func (uc *musicUsecase) GetAllMusic() ([]*domain.Music, error) {
	return uc.musicRepository.GetAll()
}

func (uc *musicUsecase) AddMusic(music *domain.Music) error {
	fmt.Println("MUSIC:::", music)
	return uc.musicRepository.Add(music)
}

func (uc *musicUsecase) GetMusicByID(id uint64) (*domain.Music, error) {
	return uc.musicRepository.GetByID(id)
}

func (uc *musicUsecase) UpdateMusic(music *domain.Music) error {
	return uc.musicRepository.Update(music)
}

func (uc *musicUsecase) DeleteMusic(id string) error {
	return uc.musicRepository.Delete(id)
}
