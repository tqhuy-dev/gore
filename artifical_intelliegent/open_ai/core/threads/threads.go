package threads

import (
	"context"
	"fmt"
	"github.com/tqhuy-dev/gore/artifical_intelliegent/open_ai/core"
)

type IOpenAIThreads interface {
	CreateThread(ctx context.Context, request *CreateThreadDto) (*CreateThreadResponseDto, error)
	RunThread(ctx context.Context, request *RunThreadRequest) (*CreateThreadResponseDto, error)
	RetrieveRunThread(ctx context.Context, request *RetrieveThreadRunRequest) (*RetrieveThreadRunResponse, error)
	SendMessageThreads(ctx context.Context, request *SendMessageThreadsRequest) (*SendMessageThreadsResponse, error)
	GetListMessagesOnThread(ctx context.Context, request *ListMessagesOnThreadRequest) (*ListMessagesOnThreadResponse, error)
}

type openAIThreads struct {
	BaseGPTOpenAI core.IGptOpenAI
}

func (thread *openAIThreads) GetListMessagesOnThread(ctx context.Context, request *ListMessagesOnThreadRequest) (*ListMessagesOnThreadResponse, error) {
	path := fmt.Sprintf("%s/v1/threads/%s/messages", core.UrlOpenAI, request.ThreadId)
	var response ListMessagesOnThreadResponse
	err := thread.BaseGPTOpenAI.GetHttpClient().Get(path, thread.BaseGPTOpenAI.GetHeaders(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (thread *openAIThreads) SendMessageThreads(ctx context.Context, request *SendMessageThreadsRequest) (*SendMessageThreadsResponse, error) {
	path := fmt.Sprintf("%s/v1/threads/%s/messages", core.UrlOpenAI, request.ThreadId)
	var response SendMessageThreadsResponse
	err := thread.BaseGPTOpenAI.GetHttpClient().Post(path, thread.BaseGPTOpenAI.GetHeaders(), request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func NewOpenAIThreads(BaseGPTOpenAI core.IGptOpenAI) IOpenAIThreads {
	return &openAIThreads{
		BaseGPTOpenAI: BaseGPTOpenAI,
	}
}

func (thread *openAIThreads) CreateThread(ctx context.Context, request *CreateThreadDto) (*CreateThreadResponseDto, error) {
	path := fmt.Sprintf("%s/v1/threads", core.UrlOpenAI)
	var response CreateThreadResponseDto
	err := thread.BaseGPTOpenAI.GetHttpClient().Post(path, thread.BaseGPTOpenAI.GetHeaders(), request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (thread *openAIThreads) RunThread(ctx context.Context, request *RunThreadRequest) (*CreateThreadResponseDto, error) {
	path := fmt.Sprintf("%s/v1/threads/%s/runs", core.UrlOpenAI, request.ThreadId)
	var response CreateThreadResponseDto
	err := thread.BaseGPTOpenAI.GetHttpClient().Post(path, thread.BaseGPTOpenAI.GetHeaders(), request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (thread *openAIThreads) RetrieveRunThread(ctx context.Context, request *RetrieveThreadRunRequest) (*RetrieveThreadRunResponse, error) {
	path := fmt.Sprintf("%s/v1/threads/%s/runs/%s", core.UrlOpenAI, request.ThreadId, request.RunId)
	var response RetrieveThreadRunResponse
	err := thread.BaseGPTOpenAI.GetHttpClient().Get(path, thread.BaseGPTOpenAI.GetHeaders(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
