package repository

import "cryptotracker/entity"

type TrackerRepository interface {
	GetWallets() []entity.Wallet
	GetWallet(id int) *entity.Wallet
	InsertWallet(wallet entity.Wallet) (*entity.Wallet, error)
	EditWallet(wallet entity.Wallet) error
	InsertTransaction(walletId int, transaction entity.Transaction) (*entity.Transaction, error)
	RemoveTransaction(walletId int, transactionId int64) error
}
