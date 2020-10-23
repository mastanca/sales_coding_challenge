package main

import (
	"log"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/cmd/web/middleware"

	"github.com/gin-gonic/gin"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/cmd/web/handlers"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/usecases"
)

func main() {
	ticketsRepository := ticket.NewInMemoryRepository()

	registerTicketSale := usecases.NewRegisterTicketSale(ticketsRepository)
	getSalesStatsPerCountry := usecases.NewGetSalesStatsPerCountryImpl(ticketsRepository)
	getUser := usecases.NewGetUserImpl()

	registerTicketSaleHandler := handlers.NewRegisterTicketSaleHandlerImpl(registerTicketSale)
	getSaleStatsPerCountryHandler := handlers.NewGetSaleStatsPerCountryHandlerImpl(getSalesStatsPerCountry)
	userLoginHandler := handlers.NewLoginHandlerImpl(getUser)

	router := gin.Default()
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/sales", middleware.AuthRequired(), registerTicketSaleHandler.Handle)
			v1.GET("/stats", middleware.AuthRequired(), getSaleStatsPerCountryHandler.Handle)
			v1.POST("/login", userLoginHandler.Handle)
		}
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatal("error initializing server")
	}
}
