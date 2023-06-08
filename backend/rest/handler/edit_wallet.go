package handler

import (
	"cryptotracker/rest/helper"
	"cryptotracker/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func EditWalletHandlerFunc(walletService *service.WalletService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletId, err := strconv.Atoi(chi.URLParam(r, "walletId"))
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		var requestPayload struct {
			Name   string `json:"name"`
			Crypto string `json:"crypto"`
			Fiat   string `json:"fiat"`
		}

		err = helper.ReadJSON(w, r, &requestPayload)
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		err = walletService.EditWallet(walletId, requestPayload.Name, requestPayload.Crypto, requestPayload.Fiat)
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		_ = helper.WriteJSON(w, http.StatusOK, nil)
	}
}
