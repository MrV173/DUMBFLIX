package repository

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type TvRepository interface {
	FindTvs() ([]models.Tv, error)
	GetTv(ID int) (models.Tv, error)
	CreateTv(tv models.Tv) (models.Tv, error)
	UpdateTv(tv models.Tv) (models.Tv, error)
	DeleteTv(tv models.Tv) (models.Tv, error)
}

func RepositoryTv(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTvs() ([]models.Tv, error) {
	var tvs []models.Tv
	err := r.db.Find(&tvs).Error

	return tvs, err
}

func (r *repository) GetTv(ID int) (models.Tv, error) {
	var tv models.Tv
	err := r.db.First(&tv, ID).Error

	return tv, err
}

func (r *repository) CreateTv(tv models.Tv) (models.Tv, error) {
	err := r.db.Create(&tv).Error

	return tv, err
}

func (r *repository) UpdateTv(tv models.Tv) (models.Tv, error) {
	err := r.db.Save(&tv).Error

	return tv, err
}

func (r *repository) DeleteTv(tv models.Tv) (models.Tv, error) {
	err := r.db.Delete(&tv).Error

	return tv, err
}
