package providers

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/vault-client-go"
	"log"
	"os"
	"time"
)

type IVault interface {
	GetClient() *vault.Client
	Encrypt(ctx context.Context, key string, data string, path string) (string, error)
	Decrypt(ctx context.Context, key string, cipherText string, path string) (string, error)
	GetSecretStore(ctx context.Context, path string) (map[string]interface{}, error)
	EncryptWithSalt(ctx context.Context, data string, path string, salt string) (string, error)
	DecryptWithSalt(ctx context.Context, cipherText string, path string, salt string) (string, error)
}
type VaultConfig struct {
	Address string
	Port    int
	Token   string
}

type vaultProvider struct {
	vault *vault.Client
}

func (v *vaultProvider) EncryptWithSalt(ctx context.Context, data string, path string, salt string) (string, error) {
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	saltStr := base64.StdEncoding.EncodeToString([]byte(salt))
	payload := map[string]interface{}{
		"plaintext": sEnc, // Dữ liệu phải được base64 encode
		"context":   saltStr,
	}
	secret, err := v.vault.Write(ctx, path, payload)
	if err != nil {
		return "", fmt.Errorf("không thể gửi yêu cầu mã hóa đến Vault: %w", err)
	}
	// Lấy ciphertext từ phản hồi
	if secret == nil || secret.Data == nil || secret.Data["ciphertext"] == nil {
		return "", fmt.Errorf("không nhận được ciphertext từ Vault")
	}
	cipherText, ok := secret.Data["ciphertext"].(string)
	if !ok {
		return "", fmt.Errorf("ciphertext không phải là chuỗi")
	}
	return cipherText, nil
}

func (v *vaultProvider) DecryptWithSalt(ctx context.Context, cipherText string, path string, salt string) (string, error) {

	sSalt := base64.StdEncoding.EncodeToString([]byte(salt))

	data := map[string]interface{}{
		"ciphertext": cipherText,
		"context":    sSalt,
	}
	secret, err := v.vault.Write(ctx, path, data)
	if err != nil {
		return "", fmt.Errorf("không thể gửi yêu cầu giải mã đến Vault: %w", err)
	}
	// Lấy plaintext đã giải mã từ phản hồi (nó đã được base64 encode)
	if secret == nil || secret.Data == nil || secret.Data["plaintext"] == nil {
		return "", fmt.Errorf("không nhận được plaintext từ Vault")
	}

	encodedPlaintext, ok := secret.Data["plaintext"].(string)
	if !ok {
		return "", fmt.Errorf("plaintext không phải là chuỗi")
	}
	// Base64 decode plaintext
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedPlaintext)
	if err != nil {
		return "", fmt.Errorf("không thể base64 decode plaintext: %w", err)
	}

	return string(decodedBytes), nil
}

func (v *vaultProvider) GetClient() *vault.Client {
	return v.vault
}

func NewVaultProvider(config VaultConfig) IVault {
	client, err := vault.New(
		vault.WithAddress(fmt.Sprintf("%s:%d", config.Address, config.Port)),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.SetToken(config.Token); err != nil {
		log.Fatal(err)
	}
	return &vaultProvider{
		vault: client,
	}
}

func (v *vaultProvider) Encrypt(ctx context.Context, key string, data string, path string) (string, error) {
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	payload := map[string]interface{}{
		"plaintext": sEnc, // Dữ liệu phải được base64 encode
	}
	pathVaultEncrypt := fmt.Sprintf("%s/%s", path, key)
	secret, err := v.vault.Write(ctx, pathVaultEncrypt, payload)
	if err != nil {
		return "", fmt.Errorf("không thể gửi yêu cầu mã hóa đến Vault: %w", err)
	}
	// Lấy ciphertext từ phản hồi
	if secret == nil || secret.Data == nil || secret.Data["ciphertext"] == nil {
		return "", fmt.Errorf("không nhận được ciphertext từ Vault")
	}
	cipherText, ok := secret.Data["ciphertext"].(string)
	if !ok {
		return "", fmt.Errorf("ciphertext không phải là chuỗi")
	}
	return cipherText, nil
}

func (v *vaultProvider) Decrypt(ctx context.Context, key string, cipherText string, path string) (string, error) {
	data := map[string]interface{}{
		"ciphertext": cipherText,
	}
	pathVaultDecrypt := fmt.Sprintf("%s/%s", path, key)
	secret, err := v.vault.Write(ctx, pathVaultDecrypt, data)
	if err != nil {
		return "", fmt.Errorf("không thể gửi yêu cầu giải mã đến Vault: %w", err)
	}
	// Lấy plaintext đã giải mã từ phản hồi (nó đã được base64 encode)
	if secret == nil || secret.Data == nil || secret.Data["plaintext"] == nil {
		return "", fmt.Errorf("không nhận được plaintext từ Vault")
	}

	encodedPlaintext, ok := secret.Data["plaintext"].(string)
	if !ok {
		return "", fmt.Errorf("plaintext không phải là chuỗi")
	}
	// Base64 decode plaintext
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedPlaintext)
	if err != nil {
		return "", fmt.Errorf("không thể base64 decode plaintext: %w", err)
	}

	return string(decodedBytes), nil
}

func (v *vaultProvider) GetSecretStore(ctx context.Context, path string) (map[string]interface{}, error) {
	secret, err := v.vault.Read(ctx, path)
	if err != nil {
		log.Fatalf("Lỗi khi đọc bí mật từ Vault tại %s: %v", path, err)
	}

	if secret == nil || secret.Data == nil {
		log.Fatalf("Không tìm thấy bí mật tại %s hoặc dữ liệu rỗng. Hãy đảm bảo bạn đã ghi bí mật vào Vault.", path)
	}

	// KV v2 lưu trữ dữ liệu thực sự trong map "data" của secret.Data
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalf("Định dạng dữ liệu không đúng từ Vault. Đảm bảo đây là KV v2.")
	}
	return data, nil
}

func loadVaultSecrets(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	secrets := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		secrets = append(secrets, line)
	}
	return secrets, scanner.Err()
}

func ReturnSecret() []string {
	secrets, err := loadVaultSecrets("/vault/secrets/database-config.txt")
	if err != nil {
		panic(err)
	}

	return secrets

}
