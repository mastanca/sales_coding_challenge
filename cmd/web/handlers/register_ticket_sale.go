package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"
	"net/http"
)

type RegisterTicketSaleHandler interface {
	Handle(c *gin.Context)
}

type registerTicketSaleHandlerImpl struct {
}

func NewRegisterTicketSaleHandlerImpl() *registerTicketSaleHandlerImpl {
	return &registerTicketSaleHandlerImpl{}
}

func (r registerTicketSaleHandlerImpl) Handle(c *gin.Context) {
	var newTicket ticket.Ticket
	if err := c.ShouldBindJSON(&newTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTicket)
}

var _ RegisterTicketSaleHandler = (*registerTicketSaleHandlerImpl)(nil)
