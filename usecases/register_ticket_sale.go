package usecases

import (
	"context"

	"github.com/pkg/errors"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"
)

type RegisterTicketSale interface {
	Execute(ctx context.Context, ticket ticket.Ticket) error
}

type registerTicketSaleImpl struct {
	repository ticket.Repository
}

func NewRegisterTicketSale(repository ticket.Repository) *registerTicketSaleImpl {
	return &registerTicketSaleImpl{repository: repository}
}

func (r registerTicketSaleImpl) Execute(ctx context.Context, ticket ticket.Ticket) error {
	if err := r.repository.Save(ctx, ticket); err != nil {
		return errors.Wrap(err, "error saving ticket")
	}
	return nil
}

var _ RegisterTicketSale = (*registerTicketSaleImpl)(nil)
