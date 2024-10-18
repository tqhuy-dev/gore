package main

import (
	"fmt"
	"github.com/s-platform/gore/crypto"
)

func main() {

	a, err := crypto.Factory(crypto.AES, "Key")
	if err != nil {
		panic(err)
	}
	r, err := a.Encrypt(crypto.EncryptCondition{PlainText: "HuyTran"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.CipherText)

	f, err := a.Decrypt(crypto.DecryptCondition{
		CipherText: r.CipherText,
		Nonce:      r.Nonce,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(f.PlainText)
}
