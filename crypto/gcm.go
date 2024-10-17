package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

type gcmCrypto struct {
	gcm cipher.AEAD
}

func (g gcmCrypto) Encrypt(condition EncryptCondition) (EncryptResult, error) {
	nonce := make([]byte, g.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return EncryptResult{}, err
	}
	cipherText := g.gcm.Seal(nil, nonce, []byte(condition.PlainText), nil)
	return EncryptResult{CipherText: string(cipherText), Nonce: nonce}, nil
}

func (g gcmCrypto) Decrypt(condition DecryptCondition) (DecryptResult, error) {
	plainText, err := g.gcm.Open(nil, condition.Nonce, []byte(condition.CipherText), nil)
	if err != nil {
		return DecryptResult{}, err
	}
	return DecryptResult{PlainText: string(plainText)}, nil
}

func NewGCMCrypto(key string) (IHandleCrypto, error) {
	key256 := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(key256[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &gcmCrypto{gcm: gcm}, nil
}
