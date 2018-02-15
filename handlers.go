package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// accounts holds all accounts
type accounts struct {
	publicKey  map[string]string
	privateKey map[string]string
}

var (
	accts *accounts
)

// AccountsRepository returns a singleton account repository
func AccountsRepository() *accounts {
	if accts == nil {
		accts = &accounts{
			publicKey:  make(map[string]string),
			privateKey: make(map[string]string),
		}
	}
	return accts
}

// Index defining Index function
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex defining TodoIndex function
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow defining TodoShow function
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}

// MakePayment - makes payment from account a to account b
func MakePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Make Payment Vars: ", vars)
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
}

// CreateNewAccount - creates a new account for the given addr
func CreateNewAccount(addr string, w http.ResponseWriter) {

	resp, err := http.Get("https://horizon-testnet.stellar.org/friendbot?addr=" + addr)
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
