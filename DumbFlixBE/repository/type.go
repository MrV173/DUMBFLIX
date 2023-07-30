package repository

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type TypeFilmRepository interface {
	FindTypeFilms() ([]models.TypeFilm, error)
	GetTypeFilm(ID int) (models.TypeFilm, error)
	CreateTypeFilm(typeFilm models.TypeFilm) (models.TypeFilm, error)
	DeleteTypeFilm(typeFilm models.TypeFilm) (models.TypeFilm, error)
}

func RepositoryTypeFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTypeFilms() ([]models.TypeFilm, error) {
	var typeFilms []models.TypeFilm
	err := r.db.Find(&typeFilms).Error

	return typeFilms, err
}

func (r *repository) GetTypeFilm(ID int) (models.TypeFilm, error) {
	var typeFilm models.TypeFilm
	err := r.db.First(&typeFilm, ID).Error

	return typeFilm, err
}

func (r *repository) CreateTypeFilm(typefilm models.TypeFilm) (models.TypeFilm, error) {
	err := r.db.Create(&typefilm).Error

	return typefilm, err
}

func (r *repository) DeleteTypeFilm(typeFilm models.TypeFilm) (models.TypeFilm, error) {
	err := r.db.Delete(&typeFilm).Error

	return typeFilm, err
}
