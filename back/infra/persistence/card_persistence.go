package persistence

import (
	"context"

	"github.com/ei-sugimoto/pokepoke/back/domain/entity"
	"github.com/ei-sugimoto/pokepoke/back/domain/repository"
	"github.com/ei-sugimoto/pokepoke/back/ent"
)

type CardPersistence struct {
	db *ent.Client
}

func NewCardPersistence(db *ent.Client) repository.CardRepository {
	return &CardPersistence{db: db}
}

// FindAll fetches all cards.
func (cp *CardPersistence) FindAll(ctx context.Context) ([]entity.Card, error) {
	cards, err := cp.db.Card.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	resCards := make([]entity.Card, len(cards))
	for _, c := range cards {
		resCards = append(resCards, entity.Card{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	return resCards, nil
}

// FindByID fetches a card by its ID.
func (cp *CardPersistence) FindByID(ctx context.Context, id int) (*entity.Card, error) {
	card, err := cp.db.Card.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Card{
		ID:   card.ID,
		Name: card.Name,
	}, nil
}

// Create creates a new card.
func (cp *CardPersistence) Create(ctx context.Context, card *entity.Card) error {
	_, err := cp.db.Card.Create().
		SetName(card.Name).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// BulkCreate creates multiple cards.
func (cp *CardPersistence) BulkCreate(ctx context.Context, cards []*entity.Card) error {
	bulk := make([]*ent.CardCreate, len(cards))
	for i, c := range cards {
		bulk[i] = cp.db.Card.Create().SetName(c.Name)
	}
	_, err := cp.db.Card.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
