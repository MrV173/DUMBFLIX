package repository

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type EpisodeTv interface {
	FindEps() ([]models.EpisodeTV, error)
	GetEps(ID int) (models.EpisodeTV, error)
	CreateEps(eps models.EpisodeTV) (models.EpisodeTV, error)
	UpdateEps(eps models.EpisodeTV) (models.EpisodeTV, error)
	DeleteEps(eps models.EpisodeTV) (models.EpisodeTV, error)
}

func RepositoryEpisodeTV(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEps() ([]models.EpisodeTV, error) {
	var epss []models.EpisodeTV
	err := r.db.Find(&epss).Error

	return epss, err
}

func (r *repository) GetEps(ID int) (models.EpisodeTV, error) {
	var eps models.EpisodeTV
	err := r.db.First(&eps, ID).Error

	return eps, err
}

func (r *repository) CreateEps(eps models.EpisodeTV) (models.EpisodeTV, error) {
	err := r.db.Create(&eps).Error

	return eps, err
}

func (r *repository) UpdateEps(eps models.EpisodeTV) (models.EpisodeTV, error) {
	err := r.db.Save(&eps).Error

	return eps, err
}

func (r *repository) DeleteEps(eps models.EpisodeTV) (models.EpisodeTV, error) {
	err := r.db.Delete(&eps).Error

	return eps, err
}
