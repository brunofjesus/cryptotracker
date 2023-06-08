package persistence

import "cryptotracker/entity"

type FsPersistence interface {
	Save(wallets []entity.Wallet) error
	Load() ([]entity.Wallet, error)
}
