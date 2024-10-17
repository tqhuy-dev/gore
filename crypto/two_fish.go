package crypto

import (
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/twofish"
	"io"
)

type twoFishCrypto struct {
	cipher *twofish.Cipher
}

func (t twoFishCrypto) Encrypt(condition EncryptCondition) (EncryptResult, error) {
	plaintTextByte := []byte(condition.PlainText)
	plaintTextByte = Pad(plaintTextByte, t.cipher.BlockSize())

	nonce := make([]byte, twofish.BlockSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return EncryptResult{}, err
	}

	// Chế độ CBC
	mode := cipher.NewCBCEncrypter(t.cipher, nonce)

	// Mã hóa
	cipherText := make([]byte, len(plaintTextByte))
	mode.CryptBlocks(cipherText, plaintTextByte)

	return EncryptResult{CipherText: string(cipherText), Nonce: nonce}, nil
}

func (t twoFishCrypto) Decrypt(condition DecryptCondition) (DecryptResult, error) {
	cipherTextByte := []byte(condition.CipherText)
	mode := cipher.NewCBCDecrypter(t.cipher, condition.Nonce)

	// Giải mã
	plainText := make([]byte, len(cipherTextByte))
	mode.CryptBlocks(plainText, cipherTextByte)

	plainText = Unpad(plainText)

	return DecryptResult{PlainText: string(plainText)}, nil
}

func NewTwoFishCrypto(key string) (IHandleCrypto, error) {
	key256 := sha256.Sum256([]byte(key))
	block, err := twofish.NewCipher(key256[:])
	if err != nil {
		return nil, err
	}
	return &twoFishCrypto{cipher: block}, nil
}
