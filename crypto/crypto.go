package crypto

import (
	"bytes"
)

type IHandleCrypto interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
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
