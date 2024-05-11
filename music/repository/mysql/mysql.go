package mysql

import (
	"musicApp/domain"

	"gorm.io/gorm"
)

type MusicRepository struct {
	db *gorm.DB
}

func NewMusicRepository(db *gorm.DB) domain.MusicRepository {
	return &MusicRepository{db}
}

func (r *MusicRepository) GetAll() ([]*domain.Music, error) {
	var musicList []*domain.Music
	if err := r.db.Find(&musicList).Error; err != nil {
		return nil, err
	}
	return musicList, nil
}

func (r *MusicRepository) Add(music *domain.Music) error {
	if err := r.db.Create(music).Error; err != nil {
		return err
	}
	return nil
}

func (r *MusicRepository) Update(music *domain.Music) error {
	if err := r.db.Save(music).Error; err != nil {
		return err
	}
	return nil
}

func (r *MusicRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Music{}).Error; err != nil {
		return err
	}
	return nil
}
