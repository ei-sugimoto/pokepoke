package cmd

import (
	"github.com/ei-sugimoto/pokepoke/back/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var _ openapi.ServerInterface = (*Server)(nil)

type Server struct {
}

func (s *Server) GetTest(ctx echo.Context) error {
	message := "health"
	res := &openapi.GetTestResponse{
		JSON200: &struct {
			Message *string "json:\"message,omitempty\""
		}{Message: &message},
	}

	return ctx.JSON(200, res.JSON200)
}

func Serve() {

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
