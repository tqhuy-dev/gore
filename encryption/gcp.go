package encryption

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
)

type IGcpSymmetricEncryption interface {
	IKMSSymmetricEncryption
	InitClient(client *kms.KeyManagementClient)
}

type gcpSymmetricEncryption struct {
	client *kms.KeyManagementClient
}

func (g *gcpSymmetricEncryption) TypeEncrypt() KMSDriverType {
	return GCPDriver
}

func (g *gcpSymmetricEncryption) InitClient(client *kms.KeyManagementClient) {
	g.client = client
}

func (g *gcpSymmetricEncryption) Encrypt(ctx context.Context, plaintext []byte, keyPath string) ([]byte, error) {
	req := &kmspb.EncryptRequest{
		Plaintext: plaintext,
		Name:      keyPath,
	}
	g.client.GetCryptoKey()
	resp, err := g.client.Encrypt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Ciphertext, nil
}

func (g *gcpSymmetricEncryption) Decrypt(ctx context.Context, ciphertext []byte, keyPath string) ([]byte, error) {
	req := &kmspb.DecryptRequest{
		Name:       keyPath,
		Ciphertext: ciphertext,
	}
	resp, err := g.client.Decrypt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Plaintext, nil
}

func NewGCPSymmetricEncryption(client *kms.KeyManagementClient) IGcpSymmetricEncryption {
	return &gcpSymmetricEncryption{
		client: client,
	}
}
