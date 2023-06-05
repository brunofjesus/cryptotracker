package handler

import (
	"cryptotracker/rest/helper"
	"cryptotracker/service"
	"net/http"
)

func GetWalletsHandlerFunc(walletService *service.WalletService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = helper.WriteJSON(w, http.StatusOK, walletService.GetWallets())
	}
}
