package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Group) {
	r := repository.RepositoryCategory(mysql.DB)
	h := handlers.CategoryHandler(r)

	e.GET("/categories", h.FindCategories)
	e.GET("/category/:id", h.GetCategory)
	e.POST("/category", h.CreateCategory)
	e.DELETE("/category/:id", h.DeleteCategory)
}
