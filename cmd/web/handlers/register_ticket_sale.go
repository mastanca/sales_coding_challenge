package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/usecases"
)

type RegisterTicketSaleHandler interface {
	Handle(c *gin.Context)
}

type registerTicketSaleHandlerImpl struct {
	registerTicketSale usecases.RegisterTicketSale
}

func NewRegisterTicketSaleHandlerImpl(registerTicketSale usecases.RegisterTicketSale) *registerTicketSaleHandlerImpl {
	return &registerTicketSaleHandlerImpl{registerTicketSale: registerTicketSale}
}

func (r registerTicketSaleHandlerImpl) Handle(c *gin.Context) {
	var newTicket ticket.Ticket
	if err := c.ShouldBindJSON(&newTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.registerTicketSale.Execute(c, newTicket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTicket)
}

var _ RegisterTicketSaleHandler = (*registerTicketSaleHandlerImpl)(nil)
