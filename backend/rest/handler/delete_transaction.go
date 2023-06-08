package handler

import (
	"cryptotracker/rest/helper"
	"cryptotracker/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func DeleteTransactionHandlerFunc(walletService *service.WalletService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletId, err := strconv.Atoi(chi.URLParam(r, "walletId"))
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		transactionId, err := strconv.ParseInt(chi.URLParam(r, "transactionId"), 10, 64)
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		err = walletService.DeleteTransaction(walletId, transactionId)
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		_ = helper.WriteJSON(w, http.StatusOK, nil)
	}
}
