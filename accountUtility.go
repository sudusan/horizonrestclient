package main

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

// GetAccountKeyPairFor returns public/private keys for name
func GetAccountKeyPairFor(name string) (string, string) {

	ar := AccountsRepository()
	pk1, ok := ar.publicKey[name]
	var puk, prk string
	if ok {
		puk = pk1
	} else {
		puk = ""
	}
	pk2, ok := ar.privateKey[name]
	if ok {
		prk = pk2
	} else {
		prk = ""
	}
	return puk, prk
}
