package routes

import (
	"tugas/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	//user

	e.POST("/books", api.MusliGGWP)
	e.GET("/bookss", api.GetAllBooks)

	return e

}
