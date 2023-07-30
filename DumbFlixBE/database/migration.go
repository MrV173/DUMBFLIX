package database

import (
	"dumbflix/models"
	"dumbflix/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(

		&models.User{},
		&models.TypeFilm{},
		&models.Category{},
		&models.Movie{},
		&models.Tv{},
		&models.Transaction{},
		&models.EpisodeTV{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration has been successfully")
}
