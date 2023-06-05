package mapper

import (
	"cryptotracker/dto"
	"cryptotracker/entity"
	"fmt"
	"strconv"
)

// MapWalletToWalletDTO converts an entity.Wallet to a dto.WalletDTO.
func MapWalletToWalletDTO(src entity.Wallet, fetchCryptoValue func(crypto string) string, dest *dto.WalletDTO) error {
	dest.Id = src.Id
	dest.Name = src.Name
	dest.Crypto = src.Crypto
	dest.Fiat = src.Fiat

	var totalCryptoAmount float64 = 0
	var totalFiatInvested float64 = 0
	for _, transaction := range src.Transactions {
		totalCryptoAmount = sumStringAndFloat(totalCryptoAmount, transaction.CryptoAmount)
		totalFiatInvested = sumStringAndFloat(totalFiatInvested, transaction.FiatInvested)
	}

	dest.CryptoUnitValue = fetchCryptoValue(dest.Crypto)
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

func sumStringAndFloat(n1 float64, n2 string) float64 {
	converted, err := strconv.ParseFloat(n2, 64)
	if err != nil {
		fmt.Printf("Cannot convert %s to float64: %v", n2, err)
		return 0
	}

	return n1 + converted
}

func subStrings(s1 string, s2 string) string {
	f1, f2, err := stringsToFloat(s1, s2)
	if err != nil {
		return "ERR"
	}
	return float64ToString(f1 - f2)
}

func mulStrings(s1 string, s2 string) string {
	f1, f2, err := stringsToFloat(s1, s2)
	if err != nil {
		return "ERR"
	}

	return float64ToString(f1 * f2)
}

func divStrings(s1 string, s2 string) string {
	f1, f2, err := stringsToFloat(s1, s2)
	if err != nil {
		return "ERR"
	}

	return float64ToString(f1 / f2)
}

func stringsToFloat(s1 string, s2 string) (float64, float64, error) {
	f1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		fmt.Printf("Cannot convert %s to float64: %v", s1, err)
		return -1, -1, err
	}

	f2, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		fmt.Printf("Cannot convert %s to float64: %v", s2, err)
		return f1, -1, err
	}

	return f1, f2, nil
}

func float64ToString(n float64) string {
	return strconv.FormatFloat(n, 'E', -1, 64)
}
