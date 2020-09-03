package db

import "github.com/jinzhu/gorm"

// Address specifies an address, its balance and the wallet it came from
type Address struct {
	gorm.Model
	WalletName   string
	AccountIndex uint64 // indicates the wallet account this is a part of
	AddressIndex uint64 // indicates the subaddress index
	BaseAddress  string // indicates the base wallet account address
	Address      string `gorm:"unique"` // this is the wallet account subaddress
	Balance      uint64
}

// Transfer is a single transfer to churn an address
type Transfer struct {
	gorm.Model
	SourceAddress      string
	DestinationAddress string
	TxHash             string
}
