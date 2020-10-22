package usecases

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"
	mocks "github.com/mastanca/SALES_MARTIN_STANCANELLI/mocks/domain/ticket"
)

func TestRegisterTicketSaleImpl_Execute(t *testing.T) {
	ctx := context.TODO()
	newTicket := ticket.Ticket{
		Country: "AR",
		Event:   "cool event",
	}
	t.Run("Success", func(t *testing.T) {
		repository := new(mocks.Repository)
		defer repository.AssertExpectations(t)
		repository.On("Save", ctx, newTicket).Return(nil)

		registerTicketSale := NewRegisterTicketSale(repository)
		err := registerTicketSale.Execute(ctx, newTicket)

		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		t.Run("repository error", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)
			repository.On("Save", ctx, newTicket).Return(errors.New("fatal"))

			registerTicketSale := NewRegisterTicketSale(repository)
			err := registerTicketSale.Execute(ctx, newTicket)

			assert.EqualError(t, err, "error saving ticket: fatal")
		})
	})
}
