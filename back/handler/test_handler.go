package handler

import (
	"github.com/ei-sugimoto/pokepoke/back/openapi"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetTest(ctx echo.Context) error {
	message := "health"
	res := &openapi.GetTestResponse{
		JSON200: &struct {
			Message *string "json:\"message,omitempty\""
		}{Message: &message},
	}

	return ctx.JSON(200, res.JSON200)
}
