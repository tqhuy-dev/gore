package main

import (
	"fmt"
	"github.com/s-platform/gore/crypto"
)

func main() {
	c, err := crypto.NewTwoFishCrypto("myverystrongpasswordo32bitlength", "1234567812345678")
	if err != nil {
		panic(err)
	}
	d, err := c.Encrypt(`HuyTran`)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	f, err := c.Decrypt(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}
