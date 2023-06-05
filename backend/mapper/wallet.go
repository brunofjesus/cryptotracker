package mapper

import (
	"cryptotracker/dto"
	"cryptotracker/entity"
)

// MapWalletToWalletDTO converts an entity.Wallet to a dto.WalletDTO.
func MapWalletToWalletDTO(src entity.Wallet, fetchCryptoValue func(crypto string) string, dest *dto.WalletDTO) error {
	dest.Id = src.Id
	dest.Name = src.Name
	dest.Crypto = src.Crypto
	dest.Fiat = src.Fiat
	dest.CryptoUnitValue = fetchCryptoValue(dest.Crypto)

	var transactions = make([]dto.TransactionDTO, 0, len(src.Transactions))
	var totalCryptoAmount float64 = 0
	var totalFiatInvested float64 = 0
	for _, transaction := range src.Transactions {
		totalCryptoAmount = sumStringAndFloat(totalCryptoAmount, transaction.CryptoAmount)
		totalFiatInvested = sumStringAndFloat(totalFiatInvested, transaction.FiatInvested)

		var transactionDto dto.TransactionDTO
		transactionDto.WalletId = src.Id
		transactionDto.TotalCryptoAmount = float64ToString(totalCryptoAmount)
		transactionDto.TotalFiatInvested = float64ToString(totalFiatInvested)
		MapTransactionToTransactionDTO(transaction, &transactionDto)

		transactions = append(transactions, transactionDto)
	}

	dest.Transactions = transactions
	dest.TotalCryptoAmount = float64ToString(totalCryptoAmount)
	dest.TotalFiatInvested = float64ToString(totalFiatInvested)

	dest.CurrentFiatValue = mulStrings(dest.CryptoUnitValue, dest.TotalCryptoAmount)

	if totalFiatInvested == 0 {
		dest.ReturnOnInvestment = "0"
		dest.ReturnOnInvestmentPercent = "100"
	} else {
		dest.ReturnOnInvestment = subStrings(dest.CurrentFiatValue, dest.TotalFiatInvested)
		dest.ReturnOnInvestmentPercent = mulStrings(dest.ReturnOnInvestment, dest.TotalFiatInvested)
	}

	if dest.CurrentFiatValue == "0" {
		dest.CryptoUnitValueBreakEven = "0"
	} else {
		dest.CryptoUnitValueBreakEven = divStrings(
			mulStrings(dest.CryptoUnitValue, dest.TotalFiatInvested),
			dest.CurrentFiatValue,
		)
	}

	return nil
}
