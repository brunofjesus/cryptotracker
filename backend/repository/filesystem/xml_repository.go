package filesystem

import (
	"cryptotracker/entity"
	"cryptotracker/repository"
	"encoding/xml"
	"fmt"
	"os"
)

type root struct {
	Wallets []entity.Wallet `xml:"wallet"`
}

type xmlRepository struct {
	filePath string
	wallets  []entity.Wallet
}

func NewRepository(filePath string) (repository.TrackerRepository, error) {
	instance := &xmlRepository{
		filePath: filePath,
		wallets:  []entity.Wallet{},
	}

	err := load(instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func load(r *xmlRepository) error {
	// Skip if the file does not exist
	if _, err := os.Stat(r.filePath); err != nil {
		return nil
	}

	contents, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	var node root
	err = xml.Unmarshal(contents, &node)
	if err != nil {
		return err
	}

	r.wallets = node.Wallets
	return nil
}

func (r *xmlRepository) GetWallets() []entity.Wallet {
	return r.wallets
}

func (r *xmlRepository) InsertWallet(wallet entity.Wallet) (*entity.Wallet, error) {
	wallet.Id = r.nextWalletId()
	r.wallets = append(r.wallets, wallet)

	err := r.saveFile()
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *xmlRepository) InsertTransaction(walletId int, transaction entity.Transaction) (*entity.Transaction, error) {
	i := r.getWalletIdx(walletId)
	if i == -1 {
		return nil, fmt.Errorf("wallet not found: %d", walletId)
	}

	transaction.Id = r.nextTransactionId(&r.wallets[i])

	transactions := r.wallets[i].Transactions
	transactions = append(transactions, transaction)
	r.wallets[i].Transactions = transactions

	err := r.saveFile()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *xmlRepository) RemoveTransaction(walletId int, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r *xmlRepository) saveFile() error {
	// backup previous file
	err := backup(r.filePath)
	if err != nil {
		return fmt.Errorf("error backing up: %v", err)
	}

	// save new file
	xmlNode := root{
		Wallets: r.wallets,
	}

	//fileData, err := xml.MarshalIndent(xmlNode, "", "\t")
	fileData, err := xml.Marshal(xmlNode)

	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, fileData, 0644)
}

func (r *xmlRepository) nextWalletId() int {
	if len(r.wallets) == 0 {
		return 1
	}
	return r.wallets[len(r.wallets)-1].Id + 1
}

func (r *xmlRepository) nextTransactionId(wallet *entity.Wallet) int64 {
	if len(wallet.Transactions) == 0 {
		return 1
	}

	return wallet.Transactions[len(wallet.Transactions)-1].Id + 1
}

func (r *xmlRepository) getWalletIdx(id int) int {
	for i, wallet := range r.wallets {
		if wallet.Id == id {
			return i
		}
	}
	return -1
}
