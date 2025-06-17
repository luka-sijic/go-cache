package server

import (

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"app/internal/router"
)

func Start() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, 
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}, 
    	AllowCredentials: false,
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.Router(e)
	e.Logger.Fatal(e.Start(":8086"))
}
