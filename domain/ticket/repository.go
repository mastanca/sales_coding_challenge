package ticket

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, ticket Ticket) error
	GetAll(ctx context.Context) (Tickets, error)
}

type inMemoryRepository struct {
	db map[string]Ticket
}

func (i inMemoryRepository) GetAll(ctx context.Context) (Tickets, error) {
	var result Tickets
	for _, v := range i.db {
		result = append(result, v)
	}
	return result, nil
}

func (i inMemoryRepository) Save(ctx context.Context, ticket Ticket) error {
	i.db[uuid.New().String()] = ticket
	return nil
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{db: make(map[string]Ticket)}
}

var _ Repository = (*inMemoryRepository)(nil)
