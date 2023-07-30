package routes

import "github.com/labstack/echo/v4"

func Routeinit(e *echo.Group) {
	UserRoutes(e)
	TypeFilmRoutes(e)
	CategoryRoutes(e)
	MovieRoutes(e)
	TvRoutes(e)
	TransactionRoutes(e)
	AuthRoutes(e)
	EpsRoutes(e)
}
