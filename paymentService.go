package main

import (
	"fmt"
	"net/http"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// GetLocalTestNetClient return a local stellar testnet
func GetLocalTestNetClient() *horizon.Client {
	newTestClient := horizon.DefaultTestNetClient
	newTestClient.URL = "https://localhost:8000/"
	return newTestClient
}

// MakeNewPayment function
func MakeNewPayment(fromPubK string, fromPrK string, toPubK string, nativeAmount string, w http.ResponseWriter) {
	fmt.Fprintln(w, "****  In MakeNewPayment function *****")
	blob := createTransaction(fromPubK, fromPrK, toPubK, nativeAmount)
	fmt.Fprintln(w, "Transaction created. blob: ", blob)
	submitTransaction(blob)
}

// CreateTransaction this creates a payment transaction and returns the blob for submitting the txn
func createTransaction(fromPubK string, fromPrK string, toPK string, nativeAmount string) string {

	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: fromPubK},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: toPK},
			b.NativeAmount{Amount: nativeAmount},
		),
	)
	if err != nil {
		panic(err)
	}
	txe, err := tx.Sign(fromPrK)
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
