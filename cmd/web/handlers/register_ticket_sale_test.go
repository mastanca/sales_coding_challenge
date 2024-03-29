package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/mock"

	mocks "github.com/mastanca/SALES_MARTIN_STANCANELLI/mocks/usecases"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/domain/ticket"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/test"
	"github.com/stretchr/testify/assert"
)

func TestRegisterTicketSaleHandlerImpl_Handle(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		registerTicketSale := new(mocks.RegisterTicketSale)
		defer registerTicketSale.AssertExpectations(t)
		registerTicketSale.On("Execute", mock.Anything, mock.Anything).Return(nil)

		handler := NewRegisterTicketSaleHandlerImpl(registerTicketSale)
		router := test.Router("/api/v1/sales", handler.Handle, http.MethodPost)

		response := test.MakeRequest(router, http.MethodPost, "/api/v1/sales", strings.NewReader(`{"country": "ar"}`))
		var responseBody ticket.Ticket
		_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, "ar", responseBody.Country)
	})

	t.Run("Error", func(t *testing.T) {
		t.Run("usecase error", func(t *testing.T) {
			registerTicketSale := new(mocks.RegisterTicketSale)
			defer registerTicketSale.AssertExpectations(t)
			registerTicketSale.On("Execute", mock.Anything, mock.Anything).Return(errors.New("fatal"))

			handler := NewRegisterTicketSaleHandlerImpl(registerTicketSale)
			router := test.Router("/api/v1/sales", handler.Handle, http.MethodPost)

			response := test.MakeRequest(router, http.MethodPost, "/api/v1/sales", strings.NewReader(`{"country": "ar"}`))
			var responseBody ticket.Ticket
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Equal(t, http.StatusInternalServerError, response.Code)
		})
		t.Run("empty body", func(t *testing.T) {
			registerTicketSale := new(mocks.RegisterTicketSale)
			defer registerTicketSale.AssertExpectations(t)

			handler := NewRegisterTicketSaleHandlerImpl(registerTicketSale)
			router := test.Router("/api/v1/sales", handler.Handle, http.MethodPost)
			response := test.MakeRequest(router, http.MethodPost, "/api/v1/sales", nil)

			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
		t.Run("invalid body", func(t *testing.T) {
			registerTicketSale := new(mocks.RegisterTicketSale)
			defer registerTicketSale.AssertExpectations(t)

			handler := NewRegisterTicketSaleHandlerImpl(registerTicketSale)
			router := test.Router("/api/v1/sales", handler.Handle, http.MethodPost)
			response := test.MakeRequest(router, http.MethodPost, "/api/v1/sales", strings.NewReader(`{"country": "ar"`))

			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
		t.Run("missing required fields in body", func(t *testing.T) {
			registerTicketSale := new(mocks.RegisterTicketSale)
			defer registerTicketSale.AssertExpectations(t)

			handler := NewRegisterTicketSaleHandlerImpl(registerTicketSale)
			router := test.Router("/api/v1/sales", handler.Handle, http.MethodPost)
			response := test.MakeRequest(router, http.MethodPost, "/api/v1/sales", strings.NewReader(`{"some": "thing"}`))

			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
	})
}
