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
const BlowFish TypeCrypto = 2
const TwoFish TypeCrypto = 3
const Chacha20 TypeCrypto = 4
const GCM TypeCrypto = 5

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

func Factory(crypto TypeCrypto, key string) (IHandleCrypto, error) {
	switch crypto {
	case AES:
		return NewAESCrypto(key)
	case BlowFish:
		return NewBlowfishCrypto(key)
	case TwoFish:
		return NewTwoFishCrypto(key)
	case GCM:
		return NewGCMCrypto(key)
	case Chacha20:
		return NewChacha20Crypto(key)
	default:
		return NewAESCrypto(key)
	}
}
