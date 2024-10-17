package crypto

import (
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/blowfish"
	"io"
)

type blowfishCrypto struct {
	cipher *blowfish.Cipher
}

func (b blowfishCrypto) Encrypt(condition EncryptCondition) (EncryptResult, error) {

	blockSize := b.cipher.BlockSize()
	paddedPlainText := Pad([]byte(condition.PlainText), blockSize)
	// Tạo nonce ngẫu nhiên cho mã hóa
	nonce := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return EncryptResult{}, err
	}

	// Khởi tạo CBC mode
	mode := cipher.NewCBCEncrypter(b.cipher, nonce)

	// Mã hóa
	cipherText := make([]byte, len(paddedPlainText))
	mode.CryptBlocks(cipherText, paddedPlainText)

	// Kết hợp IV và ciphertext để sử dụng khi giải mã
	return EncryptResult{
		CipherText: base64.StdEncoding.EncodeToString(append(nonce, cipherText...)),
		Nonce:      nonce,
	}, nil

}

func (b blowfishCrypto) Decrypt(condition DecryptCondition) (DecryptResult, error) {
	// Giải mã base64
	cipherTextBytes, err := base64.StdEncoding.DecodeString(condition.CipherText)
	if err != nil {
		return DecryptResult{}, err
	}

	// Tách IV và ciphertext
	blockSize := b.cipher.BlockSize()
	iv := cipherTextBytes[:blockSize]
	cipherTextBytes = cipherTextBytes[blockSize:]

	// Khởi tạo CBC mode
	mode := cipher.NewCBCDecrypter(b.cipher, iv)

	// Giải mã
	plainText := make([]byte, len(cipherTextBytes))
	mode.CryptBlocks(plainText, cipherTextBytes)

	return DecryptResult{PlainText: string(plainText)}, nil

}

func NewBlowfishCrypto(key string) (IHandleCrypto, error) {
	key256 := sha256.Sum256([]byte(key))
	bf, err := blowfish.NewCipher(key256[:])
	if err != nil {
		return nil, err
	}
	return &blowfishCrypto{cipher: bf}, nil
}
