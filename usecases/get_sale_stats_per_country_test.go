package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"

	mocks "github.com/mastanca/SALES_MARTIN_STANCANELLI/mocks/domain/ticket"
)

func TestGetSalesStatsPerCountryImpl_Execute(t *testing.T) {
	ctx := context.TODO()
	tickets := ticket.Tickets{{
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
		repository := new(mocks.Repository)
		defer repository.AssertExpectations(t)
		repository.On("GetAll", ctx).Return(tickets, nil)

		getSalesStatsPerCountry := NewGetSalesStatsPerCountryImpl(repository)
		statsPerCountry, err := getSalesStatsPerCountry.Execute(ctx)

		assert.NoError(t, err)
		assert.EqualValues(t, StatsPerCountry{"AR": 2, "BR": 1}, statsPerCountry)
	})
}
