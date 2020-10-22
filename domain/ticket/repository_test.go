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
