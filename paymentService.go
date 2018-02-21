package main

import (
	"fmt"
	"net/http"
	"os"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// GetLocalTestNetClient return a local stellar testnet
func GetLocalTestNetClient() *horizon.Client {

	newTestClient := horizon.DefaultTestNetClient
	//newTestClient.HomeDomainForAccount
	newTestClient.URL = os.Getenv("STELLAR_QUICKSTART_URL")
	return newTestClient
}

// MakeNewPayment function
func MakeNewPayment(fromPubK string, fromPrK string, toPubK string, nativeAmount string, w http.ResponseWriter) {
	fmt.Println("****  In MakeNewPayment function *****")
	blob := createTransaction(fromPubK, fromPrK, toPubK, nativeAmount)
	fmt.Println("Transaction created. blob: ", blob)
	submitTransaction(blob)
}

// CreateTransaction this creates a payment transaction and returns the blob for submitting the txn
func createTransaction(fromPubK string, fromPrK string, toPK string, nativeAmount string) string {

	fmt.Println("Test Network ID: ", b.TestNetwork.ID())
	fmt.Println("Test Network passphrase: ", b.TestNetwork.Passphrase)
	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: fromPubK},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: GetLocalTestNetClient()},
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

	resp, err := GetLocalTestNetClient().SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}
	fmt.Println("transaction posted in ledger:", resp.Ledger)
}
