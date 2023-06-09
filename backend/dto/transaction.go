package dto

import (
	"github.com/shopspring/decimal"
	"time"
)

type TransactionDTO struct {
	Id                     int64           `json:"id"`
	WalletId               int             `json:"walletId"`
	Time                   time.Time       `json:"time"`
	CryptoValue            decimal.Decimal `json:"cryptoValue"`
	CryptoAmount           decimal.Decimal `json:"cryptoAmount"`
	FiatInvested           decimal.Decimal `json:"fiatInvested"`
	FiatValue              decimal.Decimal `json:"fiatValue"`
	TotalCryptoAmount      decimal.Decimal `json:"totalCryptoAmount"`
	TotalFiatInvested      decimal.Decimal `json:"totalFiatInvested"`
	FiatWalletValue        decimal.Decimal `json:"fiatWalletValue"`
	FiatReturnOnInvestment decimal.Decimal `json:"fiatReturnOnInvestment"`
}
