package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	r := repository.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(r)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
