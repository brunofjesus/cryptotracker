package entity

// Wallet represents a crypto wallet
type Wallet struct {
	// Id the wallet internal id
	Id int `xml:"id"`
	// Name the Wallet's name
	Name string `xml:"name"`
	// Crypto the cryptocurrency coin ticker (e.g: BTC, ETH)
	Crypto string `xml:"crypto"`
	// Fiat the fiat currency to by used (e.g: USD, EUR)
	Fiat         string        `xml:"fiat"`
	Transactions []Transaction `xml:"transaction"`
}
