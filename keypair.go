package main

import (
	"log"

	"github.com/stellar/go/keypair"
)

// NewKeyPair generates a new private/public keypair
func NewKeyPair() (string, string) {
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}
	return pair.Seed(), pair.Address()
	//log.Println(pair.Seed())
	// SAV76USXIJOBMEQXPANUOQM6F5LIOTLPDIDVRJBFFE2MDJXG24TAPUU7
	//log.Println(pair.Address())
	// GCFXHS4GXL6BVUCXBWXGTITROWLVYXQKQLF4YH5O5JT3YZXCYPAFBJZB
}
