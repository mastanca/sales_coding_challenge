package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/test"

	mocks "github.com/mastanca/SALES_MARTIN_STANCANELLI/mocks/usecases"
	"github.com/mastanca/SALES_MARTIN_STANCANELLI/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSaleStatsPerCountryHandlerImpl_Handle(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		getSalesStatsPerCountry := new(mocks.GetSalesStatsPerCountry)
		defer getSalesStatsPerCountry.AssertExpectations(t)
		getSalesStatsPerCountry.On("Execute", mock.Anything).Return(usecases.StatsPerCountry{"AR": 5, "BO": 3}, nil)

		handler := NewGetSaleStatsPerCountryHandlerImpl(getSalesStatsPerCountry)
		router := test.Router("/api/v1/stats", handler.Handle, http.MethodGet)

		response := test.MakeRequest(router, http.MethodGet, "/api/v1/stats", nil)
		var responseBody usecases.StatsPerCountry
		_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, uint64(5), responseBody["AR"])
		assert.Equal(t, uint64(3), responseBody["BO"])
	})
}
