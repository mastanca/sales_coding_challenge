package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/cmd/web/handlers"
	"log"
)

func main() {
	registerTicketSaleHandler := handlers.NewRegisterTicketSaleHandlerImpl()

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
