package encryption

import "context"

func EncryptSample(encryptKMS IKMSSymmetricEncryption) ([]byte, error) {
	path := ""
	if encryptKMS.TypeEncrypt() == GCPDriver {
		path = "projects/gcpkms-464615/locations/global/keyRings/test-ring/cryptoKeys/test-key-2"
	} else if encryptKMS.TypeEncrypt() == VaultDriver {
		path = "transit/encrypt/my-key"
	}
	cipherText, err := encryptKMS.Encrypt(context.Background(), []byte("Avc13242"), path)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

func DecryptSample(encryptKMS IKMSSymmetricEncryption, cipherText []byte) ([]byte, error) {
	path := ""
	if encryptKMS.TypeEncrypt() == GCPDriver {
		path = "projects/gcpkms-464615/locations/global/keyRings/test-ring/cryptoKeys/test-key-2"
	} else if encryptKMS.TypeEncrypt() == VaultDriver {
		path = "transit/decrypt/my-key"
	}
	plaintext, err := encryptKMS.Decrypt(context.Background(), cipherText, path)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
