package crypto

import (
	"encoding/base64"
	"golang.org/x/crypto/chacha20"
)

type chachaCrypto struct {
	chachaCipher *chacha20.Cipher
	nonce        []byte
}

func (c chachaCrypto) Encrypt(plainText string) (string, error) {
	cipherText := make([]byte, len(plainText))
	c.chachaCipher.XORKeyStream(cipherText, []byte(plainText))

	// Kết hợp nonce và ciphertext để sử dụng khi giải mã
	return base64.StdEncoding.EncodeToString(append(c.nonce[:], cipherText...)), nil
}

func (c chachaCrypto) Decrypt(cipherText string) (string, error) {
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// Tách nonce và ciphertext
	cipherTextBytes = cipherTextBytes[12:]

	plainText := make([]byte, len(cipherTextBytes))
	c.chachaCipher.XORKeyStream(plainText, cipherTextBytes)

	return string(plainText), nil
}

func NewChaChaCrypto(key string, nonceKey string) (IHandleCrypto, error) {
	nonce := []byte(nonceKey)
	cipher, err := chacha20.NewUnauthenticatedCipher([]byte(key), nonce)
	if err != nil {
		return nil, err
	}
	return &chachaCrypto{chachaCipher: cipher, nonce: nonce}, nil
}
