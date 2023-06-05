package dto

import "time"

type TransactionDTO struct {
	Id                     int64     `json:"id"`
	WalletId               int       `json:"walletId"`
	Time                   time.Time `json:"time"`
	CryptoValue            string    `json:"cryptoValue"`
	CryptoAmount           string    `json:"cryptoAmount"`
	FiatInvested           string    `json:"fiatInvested"`
	FiatValue              string    `json:"fiatValue"`
	TotalCryptoAmount      string    `json:"totalCryptoAmount"`
	TotalFiatInvested      string    `json:"totalFiatInvested"`
	FiatWalletValue        string    `json:"fiatWalletValue"`
	FiatReturnOnInvestment string    `json:"fiatReturnOnInvestment"`
}
