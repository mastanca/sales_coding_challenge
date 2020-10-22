package ticket

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepository_Save(t *testing.T) {
	ctx := context.TODO()
	newTicket := Ticket{
		Country: "AR",
		Event:   "cooler event",
	}
	t.Run("Success", func(t *testing.T) {
		repository := NewInMemoryRepository()
		err := repository.Save(ctx, newTicket)
		assert.NoError(t, err)
	})
}

func TestInMemoryRepository_GetAll(t *testing.T) {
	ctx := context.TODO()
	tickets := Tickets{{
		Country: "AR",
		Event:   "event 1",
	}, {
		Country: "AR",
		Event:   "event 2",
	}, {
		Country: "BR",
		Event:   "event 3",
	}}

	t.Run("Success", func(t *testing.T) {
		repository := NewInMemoryRepository()
		for _, ticket := range tickets {
			_ = repository.Save(ctx, ticket)
		}
		fetchedTickets, err := repository.GetAll(ctx)

		assert.NoError(t, err)
		assert.Len(t, fetchedTickets, len(tickets))
		assert.ElementsMatch(t, fetchedTickets, tickets)
	})
}
