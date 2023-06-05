package dto

type WalletDTO struct {
	Id                        int              `json:"id"`
	Name                      string           `json:"name"`
	Crypto                    string           `json:"crypto"`
	Fiat                      string           `json:"fiat"`
	CryptoUnitValue           string           `json:"cryptoUnitValue"`
	TotalCryptoAmount         string           `json:"totalCryptoAmount"`
	TotalFiatInvested         string           `json:"totalFiatInvested"`
	CurrentFiatValue          string           `json:"currentFiatValue"`
	ReturnOnInvestment        string           `json:"returnOnInvestment"`
	ReturnOnInvestmentPercent string           `json:"returnOnInvestmentPercent"`
	CryptoUnitValueBreakEven  string           `json:"cryptoUnitValueBreakEven"`
	Transactions              []TransactionDTO `json:"transactions"`
}
