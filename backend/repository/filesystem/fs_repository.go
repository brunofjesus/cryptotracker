package filesystem

import (
	"cryptotracker/entity"
	"cryptotracker/repository"
	"cryptotracker/repository/filesystem/persistence"
	"errors"
	"fmt"
)

type fsRepository struct {
	persistence persistence.FsPersistence
	wallets     []entity.Wallet
}

func NewFsRepository(persistence persistence.FsPersistence) (repository.TrackerRepository, error) {
	wallets, err := persistence.Load()

	instance := &fsRepository{
		wallets:     wallets,
		persistence: persistence,
	}

	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *fsRepository) GetWallets() []entity.Wallet {
	return r.wallets
}

func (r *fsRepository) GetWallet(id int) *entity.Wallet {
	idx := r.getWalletIdx(id)
	if idx > -1 {
		return &r.wallets[idx]
	}
	return nil
}

func (r *fsRepository) InsertWallet(wallet entity.Wallet) (*entity.Wallet, error) {
	wallet.Id = r.nextWalletId()
	r.wallets = append(r.wallets, wallet)

	err := r.persistence.Save(r.wallets)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *fsRepository) EditWallet(wallet entity.Wallet) error {
	walletIdx := r.getWalletIdx(wallet.Id)
	if walletIdx < 0 {
		return errors.New("wallet not found")
	}

	persistedWallet := r.wallets[walletIdx]
	persistedWallet.Name = wallet.Name
	persistedWallet.Crypto = wallet.Crypto
	persistedWallet.Fiat = wallet.Fiat

	r.wallets[walletIdx] = persistedWallet
	return r.persistence.Save(r.wallets)
}

func (r *fsRepository) RemoveWallet(walletId int) error {
	walletIdx := r.getWalletIdx(walletId)
	if walletIdx < 0 {
		return errors.New("wallet not found")
	}

	r.wallets = append(r.wallets[:walletIdx], r.wallets[walletIdx+1:]...)
	return r.persistence.Save(r.wallets)
}

func (r *fsRepository) InsertTransaction(walletId int, transaction entity.Transaction) (*entity.Transaction, error) {
	i := r.getWalletIdx(walletId)
	if i < 0 {
		return nil, fmt.Errorf("wallet not found: %d", walletId)
	}

	transaction.Id = r.nextTransactionId(&r.wallets[i])

	transactions := r.wallets[i].Transactions
	transactions = append(transactions, transaction)
	r.wallets[i].Transactions = transactions

	err := r.persistence.Save(r.wallets)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *fsRepository) RemoveTransaction(walletId int, transactionId int64) error {
	i := r.getWalletIdx(walletId)
	if i < 0 {
		return fmt.Errorf("wallet not found: %d", walletId)
	}

	transactions := r.wallets[i].Transactions
	transactionIdx := -1
	for x, transaction := range transactions {
		if transaction.Id == transactionId {
			transactionIdx = x
			break
		}
	}
	if transactionIdx < 0 {
		return fmt.Errorf("transaction not found: %d", transactionId)
	}

	transactions = append(transactions[:transactionIdx], transactions[transactionIdx+1:]...)
	r.wallets[i].Transactions = transactions

	return r.persistence.Save(r.wallets)
}

func (r *fsRepository) nextWalletId() int {
	if len(r.wallets) == 0 {
		return 1
	}
	return r.wallets[len(r.wallets)-1].Id + 1
}

func (r *fsRepository) nextTransactionId(wallet *entity.Wallet) int64 {
	if len(wallet.Transactions) == 0 {
		return 1
	}

	return wallet.Transactions[len(wallet.Transactions)-1].Id + 1
}

func (r *fsRepository) getWalletIdx(id int) int {
	for i, wallet := range r.wallets {
		if wallet.Id == id {
			return i
		}
	}
	return -1
}
