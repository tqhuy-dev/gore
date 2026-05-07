package main

import (
	"context"
	"fmt"
	"github.com/tqhuy-dev/gore/providers"
	"log"
)

func main() {
	vaultProvider := providers.NewVaultProvider(providers.VaultConfig{
		Address: "http://127.0.0.1",
		Port:    8200,
		Token:   "hvs.PZ1BIYwYRVvkFkq9NHQiyye6",
	})
	//encryptDEK, err := vaultProvider.EncryptWithSalt(context.Background(), "0946515847", "transit/encrypt/my-key", "phone:3232")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("encrypted data DEK: %s\n\n", encryptDEK)
	encryptDEK := "vault:v1:IXc4KQofjDxO32qT/BtXs7el7+ES7b2voin61lD5PqDQ1UzIbxM="
	//fmt.Printf("DEK: %s\n", dekKey)
	decryptStr, err := vaultProvider.DecryptWithSalt(context.Background(), encryptDEK, "transit/decrypt/my-key", "phone:3232")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decrypt data DEK: %s\n", decryptStr)
	//dataSecretStore, err := vaultProvider.GetSecretStore(context.Background(), "/v1/secret/data/app")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(dataSecretStore)
}
