package mapper

import (
	"cryptotracker/dto"
	"cryptotracker/entity"
	"github.com/shopspring/decimal"
	"log"
)

func MapTransactionToTransactionDTO(src entity.Transaction, dest *dto.TransactionDTO) {
	var err error
	dest.Id = src.Id
	dest.Time = src.Time
	dest.CryptoValue, err = decimal.NewFromString(src.CryptoValue)
	if err != nil {
		log.Print(err)
	}

	dest.CryptoAmount, err = decimal.NewFromString(src.CryptoAmount)
	if err != nil {
		log.Print(err)
	}

	dest.FiatInvested, err = decimal.NewFromString(src.FiatInvested)
	if err != nil {
		log.Print(err)
	}

	dest.FiatValue = dest.CryptoAmount.Mul(dest.CryptoValue)

	if !dest.TotalCryptoAmount.IsZero() {
		dest.FiatWalletValue = dest.CryptoValue.Mul(dest.TotalCryptoAmount)
		dest.FiatReturnOnInvestment = dest.FiatWalletValue.Sub(dest.TotalFiatInvested)
	}
}
