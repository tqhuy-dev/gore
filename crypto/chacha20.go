package crypto

import (
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/chacha20poly1305"
	"io"
)

type chacha20Crypto struct {
	aead cipher.AEAD
}

func (c chacha20Crypto) Encrypt(condition EncryptCondition) (EncryptResult, error) {
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return EncryptResult{}, err
	}
	ciphertext := c.aead.Seal(nil, nonce, []byte(condition.PlainText), nil)
	return EncryptResult{
		CipherText: string(ciphertext),
		Nonce:      nonce,
	}, nil
}

func (c chacha20Crypto) Decrypt(condition DecryptCondition) (DecryptResult, error) {
	plaintext, err := c.aead.Open(nil, condition.Nonce, []byte(condition.CipherText), nil)
	if err != nil {
		return DecryptResult{}, nil
	}
	return DecryptResult{
		PlainText: string(plaintext),
		Nonce:     condition.Nonce,
	}, nil
}

func NewChacha20Crypto(key string) (IHandleCrypto, error) {
	key256 := sha256.Sum256([]byte(key))
	aead, err := chacha20poly1305.NewX(key256[:])
	if err != nil {
		return nil, err
	}
	return &chacha20Crypto{aead: aead}, nil
}
