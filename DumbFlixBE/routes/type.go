package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/labstack/echo/v4"
)

func TypeFilmRoutes(e *echo.Group) {
	r := repository.RepositoryTypeFilm(mysql.DB)
	h := handlers.TypeFilmHandler(r)

	e.GET("/types", h.FindTypeFilms)
	e.GET("/type/:id", h.GetTypeFilm)
	e.POST("/type", h.CreateTypeFilm)
	e.DELETE("/type/:id", h.DeleteTypeFilm)
}
