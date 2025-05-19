package utilities

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/http2"
	"io"
	"net/http"
	"time"
)

type IHttp2Client interface {
	Get(url string, headers map[string]string, response interface{}) error
	Post(url string, headers map[string]string, request interface{}, response interface{}) error
	Put(url string, request interface{}, response interface{}) error
	Delete(url string, request interface{}, response interface{}) error
}

type http2Client struct {
	client *http.Client
}

func (h *http2Client) Put(url string, request interface{}, response interface{}) error {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return BytesToStruct(body, response)
}

func (h *http2Client) Post(url string, headers map[string]string, request interface{}, response interface{}) error {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return BytesToStruct(body, response)
}

func (h *http2Client) Get(url string, headers map[string]string, response interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return BytesToStruct(body, response)

}

func (h *http2Client) Delete(url string, request interface{}, response interface{}) error {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonBody))
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return BytesToStruct(body, response)
}

func NewHttp2Client(timeoutSecond int) IHttp2Client {
	client := &http.Client{Timeout: time.Duration(timeoutSecond) * time.Second, Transport: &http2.Transport{
		AllowHTTP: true,
	}}
	return &http2Client{client: client}
}
