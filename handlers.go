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
	fmt.Println("**** In MakePayment function ****")
	fmt.Println("Make Payment Vars: ", vals)
	fmt.Println("From: ", fromPerson[0])
	fmt.Println("To: ", toPerson[0])
	fmt.Println("Amount: ", amount[0])
	fromPubK, fromPrK := GetAccountKeyPairFor(fromPerson[0])
	toPubK, _ := GetAccountKeyPairFor(toPerson[0])
	if fromPubK == "" || toPubK == "" {
		fmt.Fprintln(w, "From: ", fromPerson[0], " or To: ", toPerson[0], " does not exist")
	} else {
		fmt.Println("Ready to make a payment from: ", fromPubK, "  To: ", toPubK, " Using Private Key: ", fromPrK)
		MakeNewPayment(fromPubK, fromPrK, toPubK, amount[0], w)
	}

}

// CreateAccount defining the LoadAccount function
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "adam" {
		createAdam()
		fmt.Fprintln(w, "New account created for: ", name)
	}
	if name == "eve" {
		createEve()
		fmt.Fprintln(w, "New account created for: ", name)
	}
	/*
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
	*/
}
func createAdam() {
	adamPubKey := "GBT4SYPH3VFYEEU3BXLTZ3XKLWSXB4I7RD4IYH34NP52KI3I4ALVIUKQ"
	adamPrivKey := "SD3ONVJBHZ7G565WPQT6GQVX7PEFEGPJ7FU5YS2OAK2Q4MVA5XLOCUJI"
	ar := AccountsRepository()
	ar.publicKey["adam"] = adamPubKey
	ar.privateKey["adam"] = adamPrivKey
}
func createEve() {
	evePubKey := "GCW77MTUSGOEJFMOKPL4S27LQQE4CI4WOP2UR5NSJZLWLQP5EZL7DQTN"
	evePrivKey := "SAQYULT7KMC2GZWDVZZ7V54GBQP2SMVAYUNUVY2G77IPJDVIK5YCI3QI"
	ar := AccountsRepository()
	ar.publicKey["eve"] = evePubKey
	ar.privateKey["eve"] = evePrivKey
}

// CreateNewAccount - creates a new account for the given addr
func CreateNewAccount(addr string, w http.ResponseWriter) {
	localTestNet := GetLocalTestNetClient()
	urlForCreateAccount := localTestNet.URL + "/friendbot?addr="
	fmt.Fprintln(w, "Create URL is: ", urlForCreateAccount)
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

// GetAccount defining the GetAccount function. check if account exists locally
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
	fmt.Println("****In GetAccoutFor() method *****")
	fmt.Println("for account: ", pk)
	localTestNet := GetLocalTestNetClient()
	fmt.Println("**** retrieved the LocalTestNet *****   ", localTestNet.URL)

	acct, err := localTestNet.LoadAccount(pk)
	//acct, err := localTestNet.LoadAccount("GBT4SYPH3VFYEEU3BXLTZ3XKLWSXB4I7RD4IYH34NP52KI3I4ALVIUKQ")
	if err != nil {
		fmt.Fprintln(w, "Account does not exist for: ", pk)
		//log.Fatal(err)
	} else {
		fmt.Fprintln(w, "Found Account: ", acct.AccountID)
		fmt.Fprintln(w, "Account Balance: ", acct.Balances)
	}
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
