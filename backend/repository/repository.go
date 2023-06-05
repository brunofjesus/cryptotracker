package repository

import "cryptotracker/entity"

type TrackerRepository interface {
	GetWallets() []entity.Wallet
	InsertWallet(wallet entity.Wallet) (*entity.Wallet, error)
	InsertTransaction(walletId int, transaction entity.Transaction) (*entity.Transaction, error)
	RemoveTransaction(walletId int, id int) error
}
