package handler

import (
	"net/http"

	"github.com/ei-sugimoto/pokepoke/back/openapi"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetCards(c echo.Context) error {
	cards, err := s.cardUsec.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// 生成された型に変換
	apiCards := make([]openapi.Card, len(cards))
	for i, card := range cards {
		apiCards[i] = openapi.Card{
			Id:   &card.ID,
			Name: &card.Name,
		}
	}
	res := &openapi.GetCardsResponse{
		JSON200: &struct {
			Cards *[]openapi.Card "json:\"cards,omitempty\""
		}{Cards: &apiCards},
	}
	return c.JSON(http.StatusOK, res.JSON200)
}

func (s *Server) PostCard(c echo.Context) error {
	return nil
}
