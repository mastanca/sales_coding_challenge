package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/cmd/web/handlers"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/usecases"
)

func main() {
	ticketsRepository := ticket.NewInMemoryRepository()
	registerTicketSale := usecases.NewRegisterTicketSale(ticketsRepository)
	registerTicketSaleHandler := handlers.NewRegisterTicketSaleHandlerImpl(registerTicketSale)

	router := gin.Default()
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/sales", registerTicketSaleHandler.Handle)
			v1.GET("/stats")
		}
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatal("error initializing server")
	}
}
