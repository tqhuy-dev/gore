package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

const lenKey = 16

type aesCrypto struct {
	cipherBlock cipher.Block
}

func (a aesCrypto) Encrypt(plainText string) (string, error) {
	// Đệm cho plainText để nó có độ dài là bội số của block size
	plainTextBytes := []byte(plainText)
	plainTextBytes = Pad(plainTextBytes, aes.BlockSize)

	// Tạo vector khởi tạo (IV) ngẫu nhiên
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Mã hóa
	cipherText := make([]byte, len(plainTextBytes))
	mode := cipher.NewCBCEncrypter(a.cipherBlock, iv)
	mode.CryptBlocks(cipherText, plainTextBytes)

	// Kết hợp IV với ciphertext để sử dụng khi giải mã
	return base64.StdEncoding.EncodeToString(append(iv, cipherText...)), nil
}

func (a aesCrypto) Decrypt(cipherText string) (string, error) {
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// Tách IV và ciphertext
	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	// Giải mã
	plainTextBytes := make([]byte, len(cipherTextBytes))
	mode := cipher.NewCBCDecrypter(a.cipherBlock, iv)
	mode.CryptBlocks(plainTextBytes, cipherTextBytes)

	// Xóa đệm
	plainTextBytes = Unpad(plainTextBytes)
	return string(plainTextBytes), nil
}

func NewAESCrypto(key string) (IHandleCrypto, error) {
	if len(key) != lenKey {
		return nil, errors.New(fmt.Sprintf("len key must be equal %d", lenKey))
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &aesCrypto{cipherBlock: block}, nil
}
