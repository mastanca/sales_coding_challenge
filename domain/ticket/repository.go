package ticket

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, ticket Ticket) error
}

type inMemoryRepository struct {
	db map[string]Ticket
}

func (i inMemoryRepository) Save(ctx context.Context, ticket Ticket) error {
	i.db[uuid.New().String()] = ticket
	return nil
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{db: make(map[string]Ticket)}
}

var _ Repository = (*inMemoryRepository)(nil)
