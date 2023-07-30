package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	r := repository.RepositoryUser(mysql.DB)
	h := handlers.UserHandler(r)

	e.GET("/users", h.FindUsers)
	e.GET("/user", middleware.Auth(h.GetUser))
	e.POST("/user", h.CreateUser)
	e.PATCH("/user", middleware.Auth(middleware.UploadFile(h.UpdateUser)))
	e.DELETE("/user/:id", h.DeleteUser)
}
