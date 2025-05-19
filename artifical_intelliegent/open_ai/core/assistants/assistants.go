package assistants

import (
	"context"
	"fmt"
	"github.com/tqhuy-dev/gore/artifical_intelliegent/open_ai/core"
	"github.com/tqhuy-dev/gore/utilities"
)

type IOpenAIAssistant interface {
	CreateAssistant(ctx context.Context, request *CreateAssistantRequest) (*CreateAssistantResponse, error)
	ListAssistant(ctx context.Context, request *ListAssistantRequest) (*ListAssistantResponse, error)
}

type openAIAssistant struct {
	BaseGPTOpenAI core.IGptOpenAI
}

func NewOpenAIAssistants(BaseGPTOpenAI core.IGptOpenAI) IOpenAIAssistant {
	return &openAIAssistant{
		BaseGPTOpenAI: BaseGPTOpenAI,
	}
}

func (assistants *openAIAssistant) CreateAssistant(ctx context.Context, request *CreateAssistantRequest) (*CreateAssistantResponse, error) {
	path := fmt.Sprintf("%s/v1/assistants", core.UrlOpenAI)
	var response CreateAssistantResponse
	err := assistants.BaseGPTOpenAI.GetHttpClient().Post(path, assistants.BaseGPTOpenAI.GetHeaders(), request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (assistants *openAIAssistant) ListAssistant(ctx context.Context, request *ListAssistantRequest) (*ListAssistantResponse, error) {
	path := fmt.Sprintf("%s/v1/assistants", core.UrlOpenAI)
	buildQuery := utilities.BuildQueryUri{QueryUri: path}
	buildQuery.AddParam("limit", request.Limit)
	path = buildQuery.Build()
	var response ListAssistantResponse
	err := assistants.BaseGPTOpenAI.GetHttpClient().Get(path, assistants.BaseGPTOpenAI.GetHeaders(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
