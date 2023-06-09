package dto

import "github.com/shopspring/decimal"

type WalletDTO struct {
	Id                        int              `json:"id"`
	Name                      string           `json:"name"`
	Crypto                    string           `json:"crypto"`
	Fiat                      string           `json:"fiat"`
	CryptoUnitValue           decimal.Decimal  `json:"cryptoUnitValue"`
	TotalCryptoAmount         decimal.Decimal  `json:"totalCryptoAmount"`
	TotalFiatInvested         decimal.Decimal  `json:"totalFiatInvested"`
	CurrentFiatValue          decimal.Decimal  `json:"currentFiatValue"`
	ReturnOnInvestment        decimal.Decimal  `json:"returnOnInvestment"`
	ReturnOnInvestmentPercent decimal.Decimal  `json:"returnOnInvestmentPercent"`
	CryptoUnitValueBreakEven  decimal.Decimal  `json:"cryptoUnitValueBreakEven"`
	Transactions              []TransactionDTO `json:"transactions"`
}
