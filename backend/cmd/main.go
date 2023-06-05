package main

import (
	"cryptotracker/entity"
	"cryptotracker/repository"
	"cryptotracker/repository/filesystem"
	"cryptotracker/rest/router"
	"cryptotracker/service"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!")

	trackerRepository, err := filesystem.NewRepository("file.xml")
	if err != nil {
		panic(err)
	}

	svc := service.NewWalletService(trackerRepository)

	err = router.Start(svc)
	if err != nil {
		panic(err)
	}
}

// TODO: move to a test
func initialize(repo repository.TrackerRepository) {
	_, err := repo.InsertWallet(
		entity.Wallet{
			Id:           0,
			Name:         "Bitcoin",
			Crypto:       "BTC",
			Fiat:         "EUR",
			Transactions: nil,
		},
	)
	_, err = repo.InsertWallet(
		entity.Wallet{
			Id:           0,
			Name:         "Ethereum",
			Crypto:       "ETH",
			Fiat:         "EUR",
			Transactions: nil,
		},
	)
	_, err = repo.InsertTransaction(1, entity.Transaction{
		Id:           0,
		Time:         time.Now(),
		CryptoValue:  "25000",
		CryptoAmount: "1",
		FiatInvested: "25000",
	})
	if err != nil {
		panic(err)
		return
	}
}
