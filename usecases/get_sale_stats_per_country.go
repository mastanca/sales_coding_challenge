package usecases

import (
	"context"

	"github.com/pkg/errors"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"
)

type StatsPerCountry map[string]uint64

type GetSalesStatsPerCountry interface {
	Execute(ctx context.Context) (StatsPerCountry, error)
}

type getSalesStatsPerCountryImpl struct {
	repository ticket.Repository
}

func NewGetSalesStatsPerCountryImpl(repository ticket.Repository) *getSalesStatsPerCountryImpl {
	return &getSalesStatsPerCountryImpl{repository: repository}
}

func (g getSalesStatsPerCountryImpl) Execute(ctx context.Context) (StatsPerCountry, error) {
	soldTickets, err := g.repository.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't fetch sold tickets")
	}
	result := make(StatsPerCountry)
	for _, soldTicket := range soldTickets {
		result[soldTicket.Country] += 1
	}
	return result, nil
}

var _ GetSalesStatsPerCountry = (*getSalesStatsPerCountryImpl)(nil)
