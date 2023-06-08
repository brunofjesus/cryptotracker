package service

import (
	"cryptotracker/client/yahoo"
	"cryptotracker/dto"
	"cryptotracker/entity"
	"cryptotracker/mapper"
	"cryptotracker/repository"
	"fmt"
	"log"
	"time"
)

type WalletService struct {
	trackerRepository repository.TrackerRepository
}

func NewWalletService(trackerRepository repository.TrackerRepository) *WalletService {
	return &WalletService{
		trackerRepository: trackerRepository,
	}
}

func (w *WalletService) GetWallets() []dto.WalletDTO {
	wallets := w.trackerRepository.GetWallets()

	result := make([]dto.WalletDTO, 0, len(wallets))
	for _, wallet := range wallets {
		var walletDTO dto.WalletDTO
		err := mapper.MapWalletToWalletDTO(wallet, getQuote, &walletDTO)
		if err != nil {
			fmt.Printf("Error mapping wallet: %v", err)
		}
		result = append(result, walletDTO)
	}

	return result
}

func (w *WalletService) CreateWallet(name, crypto, fiat string) (int, error) {
	wallet, err := w.trackerRepository.InsertWallet(entity.Wallet{
		Id:           0,
		Name:         name,
		Crypto:       crypto,
		Fiat:         fiat,
		Transactions: nil,
	})

	if err != nil {
		return -1, err
	}

	return wallet.Id, nil
}

func (w *WalletService) EditWallet(id int, name, crypto, fiat string) error {
	return w.trackerRepository.EditWallet(entity.Wallet{
		Id:     id,
		Name:   name,
		Crypto: crypto,
		Fiat:   fiat,
	})
}

func (w *WalletService) CreateTransaction(
	walletId int, investedAt time.Time, cryptoValue, cryptoAmount, fiatInvested string,
) (int64, error) {

	transaction, err := w.trackerRepository.InsertTransaction(walletId, entity.Transaction{
		Id:           0,
		Time:         investedAt,
		CryptoValue:  cryptoValue,
		CryptoAmount: cryptoAmount,
		FiatInvested: fiatInvested,
	})
	if err != nil {
		return -1, err
	}

	return transaction.Id, nil
}

func getQuote(cryptoCoin string, fiat string) string {
	value, err := yahoo.GetCurrentValue(fmt.Sprintf("%s-%s", cryptoCoin, fiat))
	if err != nil {
		log.Print(err)
		return "-1"
	}
	return value
}
