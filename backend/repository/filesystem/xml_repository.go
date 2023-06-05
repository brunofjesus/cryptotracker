package filesystem

import (
	"cryptotracker/entity"
	"cryptotracker/repository"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
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

	err := instance.load()
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *xmlRepository) load() error {
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
	err := r.backupFile()
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

func (r *xmlRepository) backupFile() error {
	// Skip if the file does not exist
	if _, err := os.Stat(r.filePath); err != nil {
		return nil
	}

	// Open source for reading
	src, err := os.Open(r.filePath)
	if err != nil {
		return err
	}
	defer src.Close()

	// Create new destination file
	version := time.Now().Format("20060102T150405")
	backupFilename := fmt.Sprintf("%s.%s", r.filePath, version)

	// if destination file exists generate a new one
	if _, err := os.Stat(backupFilename); err == nil {
		// filename failed, try to find a new one
		for i := 1; i < 11; i++ {
			fileName := fmt.Sprintf("%s-%d", backupFilename, i)
			_, err = os.Stat(fileName)
			if err != nil {
				backupFilename = fileName
				break
			}
		}
	}

	des, err := os.Create(backupFilename)
	if err != nil {
		return err
	}
	defer des.Close()

	// Back it up
	_, err = io.Copy(des, src)
	if err != nil {
		return err
	}

	return des.Sync()
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
