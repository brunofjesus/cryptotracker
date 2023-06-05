package service

import (
	"cryptotracker/dto"
	"cryptotracker/entity"
	"cryptotracker/mapper"
	"cryptotracker/repository"
	"fmt"
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

func getQuote(crypto string) string {
	return "100"
}
