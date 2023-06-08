package router

import (
	"cryptotracker/rest/handler"
	ownmiddleware "cryptotracker/rest/middleware"
	"cryptotracker/service"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Start(walletService *service.WalletService) error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(ownmiddleware.Cors)
	r.Use(ownmiddleware.Hearthbeat)

	r.Get("/wallet", handler.GetWalletsHandlerFunc(walletService))
	r.Post("/wallet", handler.CreateWalletHandlerFunc(walletService))
	r.Put("/wallet/{walletId}", handler.EditWalletHandlerFunc(walletService))
	r.Delete("/wallet/{walletId}", handler.DeleteWalletHandlerFunc(walletService))
	r.Post("/wallet/{walletId}/transaction", handler.CreateTransactionHandlerFunc(walletService))
	r.Delete("/wallet/{walletId}/transaction/{transactionId}", handler.DeleteTransactionHandlerFunc(walletService))

	r.Handle("/*", handler.ServeFrontend())

	// Start the webserver
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: r,
	}

	return srv.ListenAndServe()
}
