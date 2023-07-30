package repository

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	FindMovies() ([]models.Movie, error)
	GetMovie(ID int) (models.Movie, error)
	CreateMovie(movie models.Movie) (models.Movie, error)
	UpdateMovie(movie models.Movie) (models.Movie, error)
	DeleteMovie(movie models.Movie) (models.Movie, error)
}

func RepositoryMovie(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindMovies() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Find(&movies).Error

	return movies, err
}

func (r *repository) GetMovie(ID int) (models.Movie, error) {
	var movie models.Movie
	err := r.db.First(&movie, ID).Error

	return movie, err
}

func (r *repository) CreateMovie(movie models.Movie) (models.Movie, error) {
	err := r.db.Create(&movie).Error

	return movie, err
}

func (r *repository) UpdateMovie(movie models.Movie) (models.Movie, error) {
	err := r.db.Save(&movie).Error

	return movie, err
}

func (r *repository) DeleteMovie(movie models.Movie) (models.Movie, error) {
	err := r.db.Delete(&movie).Error

	return movie, err
}
