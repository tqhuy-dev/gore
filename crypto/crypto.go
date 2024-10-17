package crypto

import (
	"bytes"
)

type EncryptResult struct {
	CipherText string
	Nonce      []byte
}

type EncryptCondition struct {
	PlainText string
}

type DecryptResult struct {
	PlainText string
	Nonce     []byte
}

type DecryptCondition struct {
	CipherText string
	Nonce      []byte
}

type IHandleCrypto interface {
	Encrypt(condition EncryptCondition) (EncryptResult, error)
	Decrypt(condition DecryptCondition) (DecryptResult, error)
}

type TypeCrypto int

const AES TypeCrypto = 1

// Hàm để đệm dữ liệu
func Pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// Hàm để xóa đệm
func Unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
