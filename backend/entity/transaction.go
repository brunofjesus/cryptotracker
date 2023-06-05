package entity

import (
	"time"
)

// Transaction represents a transaction in a wallet.
// Buy or Sell.
type Transaction struct {
	// Id The transaction id
	Id int64 `xml:"id"`
	// Time The datetime when the transaction took place
	Time time.Time `xml:"time"`
	// CryptoValue How much was a coin worth at that Time
	CryptoValue string `xml:"crypto_value"`
	// CryptoAmount The amount of crypto transacted
	CryptoAmount string `xml:"crypto_amount"`
	// FiatInvested The amount of fiat currency involved
	FiatInvested string `xml:"fiat_invested"`
}
