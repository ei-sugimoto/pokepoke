package main

import (
	"fmt"

	"github.com/ei-sugimoto/pokepoke/back/cmd"
	"github.com/ei-sugimoto/pokepoke/back/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var _ openapi.ServerInterface = (*Server)(nil)

type Server struct {
}

func (s *Server) GetTest(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{"message": "health"})
}

func main() {
	fmt.Println("hello, world")

	cmd.Migration()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://front:3000"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	server := &Server{}
	openapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8000"))
}
