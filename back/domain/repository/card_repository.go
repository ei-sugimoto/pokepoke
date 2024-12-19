package repository

import (
	"context"

	"github.com/ei-sugimoto/pokepoke/back/domain/entity"
)

type CardRepository interface {
	FindAll(ctx context.Context) ([]entity.Card, error)
	FindByID(ctx context.Context, id int) (*entity.Card, error)
	Create(ctx context.Context, card *entity.Card) error
	BulkCreate(ctx context.Context, cards []*entity.Card) error
}
