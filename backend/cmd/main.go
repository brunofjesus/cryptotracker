package main

import (
	"cryptotracker/entity"
	"cryptotracker/repository"
	"cryptotracker/repository/filesystem"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!")

	repo, err := filesystem.NewRepository("file.xml")
	if err != nil {
		panic(err)
	}

	wallets := repo.GetWallets()

	for _, wallet := range wallets {
		fmt.Printf("%d - %s (%d)\n", wallet.Id, wallet.Name, len(wallet.Transactions))
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
