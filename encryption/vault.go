package encryption

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/tqhuy-dev/gore/utilities"
	"os"
)

const vaultPathEncryption = "v1/transit/encrypt"

type IVaultEncryption interface {
	IKMSSymmetricEncryption
}

type vaultSymmetricEncryption struct {
	vaultClient *api.Client
}

func (v *vaultSymmetricEncryption) TypeEncrypt() KMSDriverType {
	return VaultDriver
}

type VaultEncryptRequest struct {
	PlainText []byte `json:"plaintext"`
}

func (v *vaultSymmetricEncryption) Encrypt(ctx context.Context, plaintext []byte, keyPath string) ([]byte, error) {

	vaultToken, err := os.ReadFile("/vault/secrets/token")
	if err = utilities.LogicalError(err, vaultToken == nil, errors.New("token not found")); err != nil {
		return nil, err
	}

	v.vaultClient.SetToken(string(vaultToken))
	base64Text := base64.StdEncoding.EncodeToString(plaintext)
	data := map[string]interface{}{
		"plaintext": base64Text,
	}
	resp, err := v.vaultClient.Logical().Write(keyPath, data)
	if err = utilities.LogicalError(err, resp == nil, errors.New("invalid request")); err != nil {
		return nil, err
	}
	ciphertext := resp.Data["ciphertext"].(string)
	return []byte(ciphertext), nil
}

func (v *vaultSymmetricEncryption) Decrypt(ctx context.Context, ciphertext []byte, keyPath string) ([]byte, error) {
	vaultToken, err := os.ReadFile("/vault/secrets/token")
	if err = utilities.LogicalError(err, vaultToken == nil, errors.New("token not found")); err != nil {
		return nil, err
	}

	v.vaultClient.SetToken(string(vaultToken))
	resp, err := v.vaultClient.Logical().Write(keyPath, map[string]interface{}{
		"ciphertext": string(ciphertext),
	})
	if err = utilities.LogicalError(err, resp == nil, errors.New("invalid response from Vault")); err != nil {
		return nil, err
	}
	plaintextB64, ok := resp.Data["plaintext"].(string)
	if !ok {
		return nil, errors.New("plaintext not found or invalid")
	}
	plaintext, err := base64.StdEncoding.DecodeString(plaintextB64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 plaintext: %w", err)
	}

	return plaintext, nil
}

func NewVaultSymmetricEncryption(
	vaultClient *api.Client) (IKMSSymmetricEncryption, error) {
	//conf := api.DefaultConfig()
	//err := conf.ConfigureTLS(&api.TLSConfig{
	//	CACert: "/run/secrets/kubernetes.io/serviceaccount/ca.crt",
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//client, err := api.NewClient(conf)
	//if err != nil {
	//	return nil, err
	//}
	return &vaultSymmetricEncryption{
		vaultClient: vaultClient,
	}, nil
}

func NewVaultClient() (*api.Client, error) {
	conf := api.DefaultConfig()
	err := conf.ConfigureTLS(&api.TLSConfig{
		CACert: "/run/secrets/kubernetes.io/serviceaccount/ca.crt",
	})
	if err != nil {
		return nil, err
	}
	conf.Address = "https://vault.vault.svc:8200"

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return client, nil
}
