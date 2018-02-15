package main

import (
	"fmt"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// MakeNewPayment function
func MakeNewPayment(from string, to string, nativeAmount string) {
	blob := createTransaction(from, to, nativeAmount)
	submitTransaction(blob)
}

// CreateTransaction this creates a payment transaction and returns the blob for submitting the txn
func createTransaction(from string, to string, nativeAmount string) string {

	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			b.NativeAmount{Amount: nativeAmount},
		),
	)
	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(from)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()

	if err != nil {
		panic(err)
	}
	return txeB64
}

func submitTransaction(blob string) {
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}
	fmt.Println("transaction posted in ledger:", resp.Ledger)
}
