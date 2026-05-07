package sample_app

import (
	"bytes"
	kms "cloud.google.com/go/kms/apiv1"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tqhuy-dev/gore/encryption"
	"github.com/tqhuy-dev/gore/providers"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func RunSampleApp() {
	var isReadiness = true
	var isLiveness = true
	e := echo.New()
	clientGcp, err := kms.NewKeyManagementClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	vaultClient, err := encryption.NewVaultClient()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		clientGcp.Close()
	}()
	e.GET("/hello", func(c echo.Context) error {
		log.Info("Hello World 2")
		return c.String(http.StatusOK, "Hello, World! 2")
	})
	e.GET("/config-map", func(c echo.Context) error {
		env := os.Getenv("APP_ENV")
		version := os.Getenv("APP_VERSION")
		log.Info("Config Map")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"env":     env,
			"version": version,
		})
	})
	e.GET("/secret", func(c echo.Context) error {
		password := os.Getenv("PASSWORD")
		log.Info("Secret")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"password": password,
		})
	})
	e.GET("/readiness", func(c echo.Context) error {
		if !isReadiness {
			return c.String(http.StatusForbidden, "I'm not ready yet")
		}
		return c.String(http.StatusOK, "I'm ready")
	})
	e.GET("/tls", func(c echo.Context) error {
		tls := c.Request().TLS
		if tls == nil || len(tls.PeerCertificates) == 0 {
			return echo.NewHTTPError(http.StatusForbidden, "No client certificate")
		}
		cert := tls.PeerCertificates[0]
		data := make([]interface{}, 0)
		for _, uri := range cert.URIs {
			data = append(data, uri)
		}
		return c.JSON(http.StatusOK, data)
	})
	e.GET("/off-readiness", func(c echo.Context) error {
		isReadiness = false
		return c.String(http.StatusOK, "Oke!")
	})
	e.GET("/on-readiness", func(c echo.Context) error {
		isReadiness = true
		return c.String(http.StatusOK, "Oke!")
	})
	e.GET("/liveness", func(c echo.Context) error {
		if !isLiveness {
			return c.String(http.StatusForbidden, "I'm not ready yet")
		}
		return c.String(http.StatusOK, "Oke!")
	})
	e.GET("/off-liveness", func(c echo.Context) error {
		isLiveness = false
		return c.String(http.StatusOK, "Oke!")
	})
	e.GET("/vault-secrets", func(c echo.Context) error {
		secrets := providers.ReturnSecret()
		return c.JSON(http.StatusOK, secrets)
	})
	e.GET("/transit-gcp", func(c echo.Context) error {
		gcpSymmetric := encryption.NewGCPSymmetricEncryption(clientGcp)
		data, err := encryption.EncryptSample(gcpSymmetric)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		plaintext, err := encryption.DecryptSample(gcpSymmetric, data)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"plaintext": string(plaintext),
		})
	})
	e.GET("/transit-vault", func(c echo.Context) error {
		vaultSymmetric, err := encryption.NewVaultSymmetricEncryption(vaultClient)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		data, err := encryption.EncryptSample(vaultSymmetric)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		plaintext, err := encryption.DecryptSample(vaultSymmetric, data)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"plaintext": string(plaintext),
		})
	})
	e.GET("/run-encrypted-data", func(c echo.Context) error {
		loadedPrivateKey, err := encryption.ImportPrivateKeyFromPEM("./private.pem")
		if err != nil {
			fmt.Printf("Lỗi tải khóa riêng: %v\n", err)
			return err
		}

		loadedPublicKey, err := encryption.ImportPublicKeyFromPEM("./public.pem")
		if err != nil {
			fmt.Printf("Lỗi tải khóa công khai: %v\n", err)
			return err
		}
		type tmp struct {
			Data string
		}
		t := tmp{
			Data: "Đây là thông điệp bí mật mà tôi muốn mã hóa bằng RSA!",
		}

		jsonData, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}
		encryptedData, err := encryption.EncryptWithPublicKey(loadedPublicKey, jsonData, "hello")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Encrypted Data: %s\n", encryptedData)

		decryptData, err := encryption.DecryptWithPrivateKey(loadedPrivateKey, encryptedData, "hello")
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, string(decryptData))
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func RunCronJob() {
	fmt.Println("Start Cronjob - Hello World\n ")
	time.Sleep(time.Second * 7)
	fmt.Println("End Cronjob - Hello World\n ")
}

func RunEncryptVault() error {
	vaultToken, _ := os.ReadFile("/vault/secrets/token")
	client := createVaultClient()
	data := map[string]interface{}{
		"plaintext": "SGVsbG8gVGFv", // "Hello Tao" base64 encoded
	}
	payload, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "https://vault.vault.svc:8200/v1/transit/encrypt/my-encryption-key", bytes.NewBuffer(payload))
	req.Header.Set("X-Vault-Token", string(vaultToken))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	dataBody := resp.Body
	body, _ := ioutil.ReadAll(dataBody)
	fmt.Println(string(body))
	defer resp.Body.Close()
	return nil
}

func createVaultClient() *http.Client {
	// Load Kubernetes CA cert
	caCert, err := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/ca.crt")
	if err != nil {
		log.Fatalf("failed to read CA cert: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("failed to append CA cert")
	}

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
}
