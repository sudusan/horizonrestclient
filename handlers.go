package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// MakePayment - makes payment from account a to account b
func MakePayment(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	fromPerson := vals["from"]
	toPerson := vals["to"]
	amount := vals["amount"]

	fmt.Fprintln(w, "Make Payment Vars: ", vals)
	fmt.Fprintln(w, "From: ", fromPerson[0])
	fmt.Fprintln(w, "To: ", toPerson[0])
	fmt.Fprintln(w, "Amount: ", amount[0])
	fromPubK, fromPrK := GetAccountKeyPairFor(fromPerson[0])
	toPubK, _ := GetAccountKeyPairFor(toPerson[0])
	if fromPubK == "" || toPubK == "" {
		fmt.Fprintln(w, "From: ", fromPerson[0], " or To: ", toPerson[0], " does not exist")
	} else {
		fmt.Fprintln(w, "Ready to make a payment from: ", fromPubK, "  To: ", toPubK, " Using Private Key: ", fromPrK)
		MakeNewPayment(fromPubK, fromPrK, toPubK, amount[0], w)
	}

}

// CreateAccount defining the LoadAccount function
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	ar := AccountsRepository()
	pk, ok := ar.publicKey[name]
	if !ok {
		pk1, pk2 := NewKeyPair()
		ar.privateKey[name] = pk1
		ar.publicKey[name] = pk2
		CreateNewAccount(pk2, w)
		fmt.Fprintln(w, "New account created for: ", name)
		fmt.Fprintln(w, "Account Key is: ", pk2)
	} else {
		fmt.Fprintln(w, "Account already exists for: ", name)
		fmt.Fprintln(w, "Account Key is: ", pk)
	}
}

// GetAccount defining the LoadAccount function
func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	ar := AccountsRepository()
	pk, ok := ar.publicKey[name]
	if !ok {
		fmt.Fprintln(w, "Account does not exist for: ", name)
	} else {
		GetAccountFor(pk, w)
	}
}

// GetAccountFor - get specific account details
func GetAccountFor(pk string, w http.ResponseWriter) {
	localTestNet := GetLocalTestNetClient()

	acct, err := localTestNet.LoadAccount(pk)
	if err != nil {
		fmt.Fprintln(w, "Account does not exist for: ", pk)
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Found Account: ", acct.AccountID)
	fmt.Fprintln(w, "Account Balance: ", acct.Balances)
	/*
		resp, err := http.Get("https://horizon-testnet.stellar.org/accounts/" + pk)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(w, "https://horizon-testnet.stellar.org/accounts/"+pk)
		fmt.Fprintln(w, string(body))
	*/
}

// CreateNewAccount - creates a new account for the given addr
func CreateNewAccount(addr string, w http.ResponseWriter) {
	localTestNet := GetLocalTestNetClient()
	urlForCreateAccount := localTestNet.URL + "friendbot?addr="
	resp, err := http.Get(urlForCreateAccount + addr)
	//resp, err := http.Get("https://horizon-testnet.stellar.org/friendbot?addr=" + addr)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(body))
}
