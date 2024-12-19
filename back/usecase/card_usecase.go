package usecase

import (
	"context"

	"github.com/ei-sugimoto/pokepoke/back/domain/entity"
	"github.com/ei-sugimoto/pokepoke/back/domain/repository"
)

type CardUsecase struct {
	cardRepo repository.CardRepository
}

func NewCardUsecase(cardRepo repository.CardRepository) *CardUsecase {
	return &CardUsecase{cardRepo: cardRepo}
}

func (cu *CardUsecase) FindAll(ctx context.Context) ([]entity.Card, error) {
	return cu.cardRepo.FindAll(ctx)
}

func (cu *CardUsecase) FindByID(ctx context.Context, id int) (*entity.Card, error) {
	return cu.cardRepo.FindByID(ctx, id)
}

func (cu *CardUsecase) Create(ctx context.Context, card *entity.Card) error {
	return cu.cardRepo.Create(ctx, card)
}

func (cu *CardUsecase) BulkCreate(ctx context.Context, cards []*entity.Card) error {
	return cu.cardRepo.BulkCreate(ctx, cards)
}
