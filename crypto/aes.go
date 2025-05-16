package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

type aesCrypto struct {
	cipherBlock cipher.Block
	nonce       []byte
}

func (a aesCrypto) Encrypt(condition EncryptCondition) (EncryptResult, error) {
	// Đệm cho plainText để nó có độ dài là bội số của block size
	plainTextBytes := []byte(condition.PlainText)
	plainTextBytes = Pad(plainTextBytes, a.cipherBlock.BlockSize())

	// Tạo vector khởi tạo (IV) ngẫu nhiên
	if a.nonce == nil {
		a.nonce = make([]byte, aes.BlockSize)
	}

	if _, err := io.ReadFull(rand.Reader, a.nonce); err != nil {
		return EncryptResult{}, err
	}

	// Mã hóa
	cipherText := make([]byte, len(plainTextBytes))
	mode := cipher.NewCBCEncrypter(a.cipherBlock, a.nonce)
	mode.CryptBlocks(cipherText, plainTextBytes)

	// Kết hợp IV với ciphertext để sử dụng khi giải mã
	return EncryptResult{
		CipherText: base64.StdEncoding.EncodeToString(append(a.nonce, cipherText...)),
		Nonce:      a.nonce,
	}, nil
}

func (a aesCrypto) Decrypt(condition DecryptCondition) (DecryptResult, error) {
	cipherTextBytes, err := base64.StdEncoding.DecodeString(condition.CipherText)
	if err != nil {
		return DecryptResult{}, err
	}

	// Tách IV và ciphertext
	cipherTextBytes = cipherTextBytes[a.cipherBlock.BlockSize():]

	// Giải mã
	plainTextBytes := make([]byte, len(cipherTextBytes))
	mode := cipher.NewCBCDecrypter(a.cipherBlock, a.nonce)
	mode.CryptBlocks(plainTextBytes, cipherTextBytes)

	// Xóa đệm
	plainTextBytes = Unpad(plainTextBytes)
	return DecryptResult{PlainText: string(plainTextBytes)}, nil
}

func NewAESCrypto(key string, nonce []byte) (IHandleCrypto, error) {
	key256 := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(key256[:])
	if err != nil {
		return nil, err
	}
	return &aesCrypto{cipherBlock: block, nonce: nonce}, nil
}
