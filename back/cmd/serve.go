package cmd

import (
	"github.com/ei-sugimoto/pokepoke/back/handler"
	"github.com/ei-sugimoto/pokepoke/back/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://front:3000"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	server := handler.NewServer()
	openapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8000"))

}
