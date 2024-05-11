package usecase

import (
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
	// You can implement any additional logic here, like validation
	return uc.musicRepository.Add(music)
}

func (uc *musicUsecase) UpdateMusic(music *domain.Music) error {
	// You can implement any additional logic here, like validation
	return uc.musicRepository.Update(music)
}

func (uc *musicUsecase) DeleteMusic(id string) error {
	return uc.musicRepository.Delete(id)
}
