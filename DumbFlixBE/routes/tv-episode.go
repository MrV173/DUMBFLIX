package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/labstack/echo/v4"
)

func EpsRoutes(e *echo.Group) {
	r := repository.RepositoryEpisodeTV(mysql.DB)
	h := handlers.EpsHandler(r)

	e.GET("/episodes", h.FindEps)
	e.GET("/episode/:id", h.GetEps)
	e.POST("/episode", middleware.UploadFile(h.CreateEps))
}
