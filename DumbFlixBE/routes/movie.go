package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/labstack/echo/v4"
)

func MovieRoutes(e *echo.Group) {
	r := repository.RepositoryMovie(mysql.DB)
	h := handlers.MovieHandler(r)

	e.GET("/movies", h.FindMovies)
	e.GET("/movie/:id", h.GetMovie)
	e.POST("/movie", middleware.UploadFile(h.CreateMovie))
	e.PATCH("/movie/:id", h.UpdateMovie)
	e.DELETE("/movie/:id", h.DeleteMovie)
}
