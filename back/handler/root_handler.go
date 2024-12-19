package handler

import (
	"github.com/ei-sugimoto/pokepoke/back/infra"
	"github.com/ei-sugimoto/pokepoke/back/infra/persistence"
	"github.com/ei-sugimoto/pokepoke/back/openapi"
	"github.com/ei-sugimoto/pokepoke/back/usecase"
)

var _ openapi.ServerInterface = (*Server)(nil)

type Server struct {
	cardUsec *usecase.CardUsecase
}

func NewServer() *Server {
	db := infra.Conn()
	cardRepo := persistence.NewCardPersistence(db)
	cardUsec := usecase.NewCardUsecase(cardRepo)
	return &Server{cardUsec: cardUsec}
}
