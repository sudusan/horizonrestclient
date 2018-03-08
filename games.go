package main

import (
	"fmt"

	b "github.com/stellar/go/build"
)

type prizeDesk struct {
	seed    string
	secret  string
	assetID string
}

type gameAccount struct {
	seed   string
	secret string
}

type user1 struct {
	seed   string
	secret string
}

// CreateTrustForAsset - this will create trust for a specific asset. execute this only once
func CreateTrustForAsset(source gameAccount, assetString string, issuer string, limit string) {

	tx, err := b.Transaction(
		b.SourceAccount{source.seed},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: GetLocalTestNetClient()},
		b.Trust(assetString, issuer, b.Limit(limit)),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(source.secret)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
}

func main() {
	pd := prizeDesk{"GBT4SYPH3VFYEEU3BXLTZ3XKLWSXB4I7RD4IYH34NP52KI3I4ALVIUKQ", "SD3ONVJBHZ7G565WPQT6GQVX7PEFEGPJ7FU5YS2OAK2Q4MVA5XLOCUJI", "TKT"}
	game := gameAccount{"GCW77MTUSGOEJFMOKPL4S27LQQE4CI4WOP2UR5NSJZLWLQP5EZL7DQTN", "SAQYULT7KMC2GZWDVZZ7V54GBQP2SMVAYUNUVY2G77IPJDVIK5YCI3QI"}
	u1 := user1{"GCLWJX5YETG7PA2AOWC237Z45UQMSKIUZJ5AMH6LZCD6DALZWZHLTVHU", "SAZL2ZEQHICI4LDU6TFCTKYXYGLQTZFCXNHE4KQRZKDKPKXTQBGDXVAY"}

	CreateTrustForAsset(game, pd.assetID, pd.seed, "100000")
}
