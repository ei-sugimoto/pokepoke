package main

import (
	"fmt"

	"github.com/ei-sugimoto/pokepoke/back/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var _ openapi.ServerInterface = (*Server)(nil)

type Server struct {
}

func (s *Server) GetTest(ctx echo.Context) error {
	return ctx.String(200, "health")
}

func main() {
	fmt.Println("hello, world")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := &Server{}
	openapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8000"))
}
