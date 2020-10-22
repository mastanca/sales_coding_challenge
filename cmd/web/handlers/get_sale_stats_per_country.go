package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/usecases"
)

type GetSaleStatsPerCountryHandler interface {
	Handle(c *gin.Context)
}

type getSaleStatsPerCountryHandlerImpl struct {
	getSalesStatsPerCountry usecases.GetSalesStatsPerCountry
}

func NewGetSaleStatsPerCountryHandlerImpl(getSaleStatsPerCountry usecases.GetSalesStatsPerCountry) *getSaleStatsPerCountryHandlerImpl {
	return &getSaleStatsPerCountryHandlerImpl{getSalesStatsPerCountry: getSaleStatsPerCountry}
}

func (g getSaleStatsPerCountryHandlerImpl) Handle(c *gin.Context) {
	statsPerCountry, err := g.getSalesStatsPerCountry.Execute(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statsPerCountry)
}

var _ GetSaleStatsPerCountryHandler = (*getSaleStatsPerCountryHandlerImpl)(nil)
