package usecases

import (
	"context"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"

	mocks "github.com/mastanca/SALES_MARTIN_STANCANELLI/mocks/domain/ticket"
)

func TestGetSalesStatsPerCountryImpl_Execute(t *testing.T) {
	ctx := context.TODO()

	t.Run("Success", func(t *testing.T) {
		t.Run("multiple countries", func(t *testing.T) {
			t.Run("single value per country", func(t *testing.T) {
				tickets := ticket.Tickets{{
					Country: "AR",
					Event:   "event 1",
				}, {
					Country: "BR",
					Event:   "event 2",
				}}
				repository := new(mocks.Repository)
				defer repository.AssertExpectations(t)
				repository.On("GetAll", ctx).Return(tickets, nil)

				getSalesStatsPerCountry := NewGetSalesStatsPerCountryImpl(repository)
				statsPerCountry, err := getSalesStatsPerCountry.Execute(ctx)

				assert.NoError(t, err)
				assert.EqualValues(t, StatsPerCountry{"AR": 1, "BR": 1}, statsPerCountry)
			})
			t.Run("multiple values per country", func(t *testing.T) {
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

				repository := new(mocks.Repository)
				defer repository.AssertExpectations(t)
				repository.On("GetAll", ctx).Return(tickets, nil)

				getSalesStatsPerCountry := NewGetSalesStatsPerCountryImpl(repository)
				statsPerCountry, err := getSalesStatsPerCountry.Execute(ctx)

				assert.NoError(t, err)
				assert.EqualValues(t, StatsPerCountry{"AR": 2, "BR": 1}, statsPerCountry)
			})
		})
	})

	t.Run("Error", func(t *testing.T) {
		t.Run("repository error", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)
			repository.On("GetAll", ctx).Return(nil, errors.New("fatal"))

			getSalesStatsPerCountry := NewGetSalesStatsPerCountryImpl(repository)
			statsPerCountry, err := getSalesStatsPerCountry.Execute(ctx)

			assert.EqualError(t, err, "couldn't fetch sold tickets: fatal")
			assert.Nil(t, statsPerCountry)
		})
	})
}
