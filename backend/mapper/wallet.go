package mapper

import (
	"cryptotracker/dto"
	"cryptotracker/entity"
	"github.com/shopspring/decimal"
	"log"
)

// MapWalletToWalletDTO converts an entity.Wallet to a dto.WalletDTO.
func MapWalletToWalletDTO(src entity.Wallet, fetchCryptoValue func(crypto string, fiat string) string, dest *dto.WalletDTO) error {
	dest.Id = src.Id
	dest.Name = src.Name
	dest.Crypto = src.Crypto
	dest.Fiat = src.Fiat
	dest.CryptoUnitValue, _ = decimal.NewFromString(fetchCryptoValue(dest.Crypto, dest.Fiat))

	var transactions = make([]dto.TransactionDTO, 0, len(src.Transactions))
	var totalCryptoAmount = decimal.Zero
	var totalFiatInvested = decimal.Zero
	for _, transaction := range src.Transactions {
		transactionAmount, err := decimal.NewFromString(transaction.CryptoAmount)
		if err != nil {
			log.Print(err)
		}

		transactionFiatInvested, err := decimal.NewFromString(transaction.FiatInvested)
		if err != nil {
			log.Print(err)
		}

		totalCryptoAmount = totalCryptoAmount.Add(transactionAmount)
		totalFiatInvested = totalFiatInvested.Add(transactionFiatInvested)

		var transactionDto dto.TransactionDTO
		transactionDto.WalletId = src.Id
		transactionDto.TotalCryptoAmount = totalCryptoAmount
		transactionDto.TotalFiatInvested = totalFiatInvested
		MapTransactionToTransactionDTO(transaction, &transactionDto)

		transactions = append(transactions, transactionDto)
	}

	dest.Transactions = transactions
	dest.TotalCryptoAmount = totalCryptoAmount
	dest.TotalFiatInvested = totalFiatInvested

	dest.CurrentFiatValue = dest.CryptoUnitValue.Mul(dest.TotalCryptoAmount)

	if totalFiatInvested.IsZero() || totalFiatInvested.IsNegative() {
		dest.ReturnOnInvestment = decimal.Zero
		dest.ReturnOnInvestmentPercent = decimal.NewFromInt(100)
	} else {
		dest.ReturnOnInvestment = dest.CurrentFiatValue.Sub(dest.TotalFiatInvested)
		dest.ReturnOnInvestmentPercent = dest.ReturnOnInvestment.Div(dest.TotalFiatInvested).
			Mul(decimal.NewFromInt(100))
	}

	if dest.CurrentFiatValue.IsZero() {
		dest.CryptoUnitValueBreakEven = decimal.Zero
	} else {
		dest.CryptoUnitValueBreakEven = dest.CryptoUnitValue.Mul(dest.TotalFiatInvested).Div(dest.CurrentFiatValue)
	}

	return nil
}
