package handler

import (
	"cryptotracker/rest/helper"
	"cryptotracker/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

func CreateTransactionHandlerFunc(walletService *service.WalletService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletId, err := strconv.Atoi(chi.URLParam(r, "walletId"))
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		var requestPayload struct {
			Time         time.Time `json:"time"`
			CryptoValue  string    `json:"cryptoValue"`
			CryptoAmount string    `json:"cryptoAmount"`
			FiatInvested string    `json:"fiatInvested"`
		}

		err = helper.ReadJSON(w, r, &requestPayload)
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		transactionId, err := walletService.CreateTransaction(
			walletId, requestPayload.Time,
			requestPayload.CryptoValue, requestPayload.CryptoAmount,
			requestPayload.FiatInvested,
		)

		_ = helper.WriteJSON(w, http.StatusCreated, transactionId)
	}
}
