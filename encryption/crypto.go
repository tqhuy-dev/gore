package encryption

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func AES256Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	// 1. Tạo Block Cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("không thể tạo AES cipher block: %v", err)
	}

	// 2. Tạo GCM (Galois/Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("không thể tạo GCM: %v", err)
	}

	// 3. Tạo Nonce (Number used once) ngẫu nhiên
	// Nonce phải là duy nhất cho mỗi quá trình mã hóa với cùng một khóa
	// Kích thước Nonce cho GCM thường là 12 bytes
	nonce := make([]byte, aesGCM.NonceSize()) // Lấy kích thước Nonce khuyến nghị từ GCM
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("không thể tạo Nonce: %v", err)
	}

	// 4. Mã hóa dữ liệu
	// Seal mã hóa plaintext và thêm tag xác thực (authentication tag)
	// Nó trả về nonce + ciphertext + tag
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil) // Nonce được prepended vào ciphertext
	return ciphertext, nil
}

func AES256Decrypt(ciphertext []byte, key []byte) ([]byte, error) {

	// 2. Tạo Block Cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("không thể tạo AES cipher block: %v", err)
	}

	// 3. Tạo GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("không thể tạo GCM: %v", err)
	}

	// 4. Tách Nonce và phần còn lại của ciphertext
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext quá ngắn để chứa Nonce")
	}

	nonce, encryptedData := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 5. Giải mã dữ liệu
	// Open sẽ giải mã và xác thực tag. Nếu tag không khớp, nó trả về lỗi.
	plaintext, err := aesGCM.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, fmt.Errorf("không thể giải mã hoặc xác thực ciphertext: %v", err)
	}

	return plaintext, nil
}

func GenerateKey(length int) ([]byte, error) {
	dek := make([]byte, length) // 256 bits
	_, err := rand.Read(dek)
	if err != nil {
		return nil, fmt.Errorf("failed to generate DEK: %v", err)
	}
	return dek, nil
}

// EncryptWithPublicKey mã hóa dữ liệu bằng khóa công khai.
func EncryptWithPublicKey(publicKey *rsa.PublicKey, data []byte, label string) ([]byte, error) {
	// rsa.OAEPHash (Optimal Asymmetric Encryption Padding) được khuyến nghị cho bảo mật.
	// Bạn có thể chọn hash function phù hợp, ví dụ sha256.New().
	encryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, data, []byte(label))
	if err != nil {
		return nil, fmt.Errorf("không thể mã hóa dữ liệu: %w", err)
	}
	return encryptedData, nil
}

// DecryptWithPrivateKey giải mã dữ liệu bằng khóa riêng.
func DecryptWithPrivateKey(privateKey *rsa.PrivateKey, encryptedData []byte, label string) ([]byte, error) {
	// rsa.OAEPHash (Optimal Asymmetric Encryption Padding) được khuyến nghị cho bảo mật.
	// Bạn có thể chọn hash function phù hợp, ví dụ sha256.New().
	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedData, []byte(label))
	if err != nil {
		return nil, fmt.Errorf("không thể giải mã dữ liệu: %w", err)
	}
	return decryptedData, nil
}

// ImportPrivateKeyFromPEM đọc khóa riêng từ định dạng PEM.
func ImportPrivateKeyFromPEM(filename string) (*rsa.PrivateKey, error) {
	privateKeyPEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("không thể đọc file khóa riêng: %w", err)
	}
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("không phải là block RSA PRIVATE KEY hợp lệ")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("không thể phân tích cú pháp khóa riêng: %w", err)
	}
	return privateKey, nil
}

// ImportPublicKeyFromPEM đọc khóa công khai từ định dạng PEM.
func ImportPublicKeyFromPEM(filename string) (*rsa.PublicKey, error) {
	publicKeyPEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("không thể đọc file khóa công khai: %w", err)
	}
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("không phải là block RSA PUBLIC KEY hợp lệ")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("không thể phân tích cú pháp khóa công khai: %w", err)
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("không phải là khóa công khai RSA hợp lệ")
	}
	return publicKey, nil
}

// GenerateRSAKeyPair tạo một cặp khóa RSA và trả về khóa riêng và khóa công khai.
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("không thể tạo khóa riêng: %w", err)
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// ExportPrivateKeyToPEM lưu khóa riêng vào định dạng PEM.
func ExportPrivateKeyToPEM(privateKey *rsa.PrivateKey, filename string) error {
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("không thể tạo file cho khóa riêng: %w", err)
	}
	defer file.Close()
	return pem.Encode(file, privateKeyPEM)
}

// ExportPublicKeyToPEM lưu khóa công khai vào định dạng PEM.
func ExportPublicKeyToPEM(publicKey *rsa.PublicKey, filename string) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("không thể marshal khóa công khai: %w", err)
	}
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("không thể tạo file cho khóa công khai: %w", err)
	}
	defer file.Close()
	return pem.Encode(file, publicKeyPEM)
}

type IKMSSymmetricEncryption interface {
	TypeEncrypt() KMSDriverType
	Encrypt(ctx context.Context, plaintext []byte, keyPath string) ([]byte, error)
	Decrypt(ctx context.Context, ciphertext []byte, keyPath string) ([]byte, error)
}

type KMSDriverType int8

const (
	GCPDriver   KMSDriverType = 1
	VaultDriver KMSDriverType = 2
)
