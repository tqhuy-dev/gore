package crypto

import (
	"crypto/cipher"
	"golang.org/x/crypto/twofish"
)

type twoFishCrypto struct {
	cipher *twofish.Cipher
	iv     []byte
}

func (t twoFishCrypto) Encrypt(plainText string) (string, error) {
	plaintTextByte := []byte(plainText)
	plaintTextByte = Pad(plaintTextByte, t.cipher.BlockSize())

	// Chế độ CBC
	mode := cipher.NewCBCEncrypter(t.cipher, t.iv)

	// Mã hóa
	cipherText := make([]byte, len(plaintTextByte))
	mode.CryptBlocks(cipherText, plaintTextByte)

	return string(cipherText), nil
}

func (t twoFishCrypto) Decrypt(cipherText string) (string, error) {
	cipherTextByte := []byte(cipherText)
	mode := cipher.NewCBCDecrypter(t.cipher, t.iv)

	// Giải mã
	plainText := make([]byte, len(cipherTextByte))
	mode.CryptBlocks(plainText, cipherTextByte)

	plainText = Unpad(plainText)

	return string(plainText), nil
}

func NewTwoFishCrypto(key string, ivText string) (IHandleCrypto, error) {
	block, err := twofish.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &twoFishCrypto{cipher: block, iv: []byte(ivText)}, nil
}
