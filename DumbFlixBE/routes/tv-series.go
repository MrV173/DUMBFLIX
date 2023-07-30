package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/labstack/echo/v4"
)

func TvRoutes(e *echo.Group) {
	r := repository.RepositoryTv(mysql.DB)
	h := handlers.TvHandler(r)

	e.GET("/tvs", h.FindTvs)
	e.GET("/tv/:id", h.GetTv)
	e.POST("/tv", middleware.UploadFile(h.CreateTv))
	e.PATCH("/tv/:id", h.UpdateTv)
	e.DELETE("/tv/:id", h.DeleteTv)
}
