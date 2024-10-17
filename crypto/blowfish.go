package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/blowfish"
	"io"
)

type blowfishCrypto struct {
	cipher *blowfish.Cipher
}

func (b blowfishCrypto) Encrypt(plainText string) (string, error) {

	blockSize := b.cipher.BlockSize()
	padding := blockSize - len(plainText)%blockSize
	paddedPlainText := append([]byte(plainText), bytes.Repeat([]byte{byte(padding)}, padding)...)

	// Tạo nonce ngẫu nhiên cho mã hóa
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Khởi tạo CBC mode
	mode := cipher.NewCBCEncrypter(b.cipher, iv)

	// Mã hóa
	cipherText := make([]byte, len(paddedPlainText))
	mode.CryptBlocks(cipherText, paddedPlainText)

	// Kết hợp IV và ciphertext để sử dụng khi giải mã
	return base64.StdEncoding.EncodeToString(append(iv, cipherText...)), nil

}

func (b blowfishCrypto) Decrypt(cipherText string) (string, error) {
	// Giải mã base64
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// Tách IV và ciphertext
	blockSize := blowfish.BlockSize
	iv := cipherTextBytes[:blockSize]
	cipherTextBytes = cipherTextBytes[blockSize:]

	// Khởi tạo CBC mode
	mode := cipher.NewCBCDecrypter(b.cipher, iv)

	// Giải mã
	plainText := make([]byte, len(cipherTextBytes))
	mode.CryptBlocks(plainText, cipherTextBytes)
	plainText = Unpad(plainText)

	return string(plainText), nil

}

func NewBlowfishCrypto(key string) (IHandleCrypto, error) {
	bf, err := blowfish.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &blowfishCrypto{cipher: bf}, nil
}
