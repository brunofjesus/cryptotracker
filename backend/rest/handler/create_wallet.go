package handler

import (
	"cryptotracker/rest/helper"
	"cryptotracker/service"
	"net/http"
)

func CreateWalletHandlerFunc(walletService *service.WalletService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestPayload struct {
			Name   string `json:"name"`
			Crypto string `json:"crypto"`
			Fiat   string `json:"fiat"`
		}

		err := helper.ReadJSON(w, r, &requestPayload)
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		walletId, err := walletService.CreateWallet(requestPayload.Name, requestPayload.Crypto, requestPayload.Fiat)
		if err != nil {
			_ = helper.ErrorJSON(w, err, http.StatusBadGateway)
			return
		}

		_ = helper.WriteJSON(w, http.StatusCreated, walletId)
	}
}
