package crypto

import "testing"

func TestAES(t *testing.T) {
	aes, _ := NewAESCrypto("admin", nil)
	resultEncryp, _ := aes.Encrypt(EncryptCondition{
		PlainText: "123456",
	})
	nonce := resultEncryp.Nonce
	aes2, _ := NewAESCrypto("admin", nonce)
	resultEncryp2, _ := aes2.Decrypt(DecryptCondition{
		CipherText: resultEncryp.CipherText,
		Nonce:      nonce,
	})
	if resultEncryp2.PlainText != "123456" {
		t.Fail()
	}
}
