package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	r := repository.RepositoryTransaction(mysql.DB)
	h := handlers.TransactionHandler(r)

	e.GET("/transactions", h.FindTransaction)
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.POST("/notification", h.Notification)
}
