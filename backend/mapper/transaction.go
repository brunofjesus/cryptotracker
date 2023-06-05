package mapper

import (
	"cryptotracker/dto"
	"cryptotracker/entity"
)

func MapTransactionToTransactionDTO(src entity.Transaction, dest *dto.TransactionDTO) {
	dest.Id = src.Id
	dest.Time = src.Time
	dest.CryptoValue = src.CryptoValue
	dest.CryptoAmount = src.CryptoAmount
	dest.FiatInvested = src.FiatInvested

	dest.FiatValue = mulStrings(src.CryptoAmount, src.CryptoValue)
	if dest.TotalCryptoAmount != "" {
		dest.FiatWalletValue = mulStrings(dest.CryptoValue, dest.TotalCryptoAmount)
		dest.FiatReturnOnInvestment = subStrings(dest.FiatWalletValue, dest.TotalFiatInvested)
	}
}
