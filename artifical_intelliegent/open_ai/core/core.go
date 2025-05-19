package core

import (
	"fmt"
	"github.com/tqhuy-dev/gore/utilities"
)

type IGptOpenAI interface {
	GetHeaders() map[string]string
	GetKeys() string
	GetHttpClient() utilities.IHttp2Client
}

type gptOpenAI struct {
	openAPIKey   string
	IHttp2Client utilities.IHttp2Client
}

func (g *gptOpenAI) GetKeys() string {
	return g.openAPIKey
}

func (g *gptOpenAI) GetHttpClient() utilities.IHttp2Client {
	return g.IHttp2Client
}

func (g *gptOpenAI) GetHeaders() map[string]string {
	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", g.openAPIKey),
		"Content-Type":  "application/json",
		"OpenAI-Beta":   "assistants=v2",
	}
}

func NewGptOpenAI(openAPIKey string, IHttp2Client utilities.IHttp2Client) IGptOpenAI {
	return &gptOpenAI{
		openAPIKey:   openAPIKey,
		IHttp2Client: IHttp2Client,
	}
}
