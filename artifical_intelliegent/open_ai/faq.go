package open_ai

import (
	"context"
	"github.com/tqhuy-dev/gore/artifical_intelliegent/open_ai/core"
	"github.com/tqhuy-dev/gore/artifical_intelliegent/open_ai/core/assistants"
	"github.com/tqhuy-dev/gore/artifical_intelliegent/open_ai/core/threads"
	"time"
)

type IChatBot interface {
	SendMessageToBot(ctx context.Context, request *SendMessageToBotRequest) (string, error)
	RunThreadsProcess(ctx context.Context, request *RunThreadsProcessRequest) (*RunResponse, error)
	GetLatestMessages(ctx context.Context, threadId string) (*SendMessageToBotResponse, error)
}

type chatBot struct {
	IOpenAIAssistant assistants.IOpenAIAssistant
	IOpenAIThreads   threads.IOpenAIThreads
}

func NewChatBot(IGptOpenAI core.IGptOpenAI) IChatBot {
	newOpenAIAssistant := assistants.NewOpenAIAssistants(IGptOpenAI)
	newOpenAIThreads := threads.NewOpenAIThreads(IGptOpenAI)
	return &chatBot{
		IOpenAIAssistant: newOpenAIAssistant,
		IOpenAIThreads:   newOpenAIThreads,
	}
}

type SendMessageToBotRequest struct {
	ThreadId    string `json:"thread_id"`
	AssistantId string `json:"assistant_id"`
	Message     string `json:"message"`
	Role        string `json:"role"`
}

type SendMessageToBotResponse struct {
	Message string `json:"message"`
	Role    string `json:"role"`
}

type RunThreadsProcessRequest struct {
	ThreadId    string `json:"thread_id"`
	AssistantId string `json:"assistant_id"`
	TimeDelay   int    `json:"time_delay"`
}

type RunResponse struct {
	Status string `json:"status"`
	RunId  string `json:"run_id"`
}

func (chatBot *chatBot) SendMessageToBot(ctx context.Context, request *SendMessageToBotRequest) (string, error) {
	resp, err := chatBot.IOpenAIThreads.SendMessageThreads(ctx, &threads.SendMessageThreadsRequest{
		ThreadId: request.ThreadId,
		Content:  request.Message,
		Role:     request.Role,
	})
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
func (chatBot *chatBot) RunThreadsProcess(ctx context.Context, request *RunThreadsProcessRequest) (*RunResponse, error) {
	runThreads, err := chatBot.IOpenAIThreads.RunThread(ctx, &threads.RunThreadRequest{
		ThreadId:    request.ThreadId,
		AssistantId: request.AssistantId,
	})
	if err != nil {
		return nil, err
	}

	for {
		retrieveThreadRunResponse, errRetrieveRun := chatBot.IOpenAIThreads.RetrieveRunThread(ctx, &threads.RetrieveThreadRunRequest{
			ThreadId: request.ThreadId,
			RunId:    runThreads.Id,
		})
		if errRetrieveRun != nil {
			return nil, errRetrieveRun
		}
		if retrieveThreadRunResponse.Status == "completed" || retrieveThreadRunResponse.Status == "requires_action" {
			return &RunResponse{
				Status: retrieveThreadRunResponse.Status,
				RunId:  runThreads.Id,
			}, nil
		}
		time.Sleep(time.Duration(request.TimeDelay) * time.Millisecond)
	}
}

func (chatBot *chatBot) GetLatestMessages(ctx context.Context, threadId string) (*SendMessageToBotResponse, error) {
	messagesOnThreadResponse, err := chatBot.IOpenAIThreads.GetListMessagesOnThread(ctx, &threads.ListMessagesOnThreadRequest{
		ThreadId: threadId,
	})
	if err != nil {
		return nil, err
	}
	if len(messagesOnThreadResponse.Data) > 0 && len(messagesOnThreadResponse.Data[0].Content) > 0 {
		return &SendMessageToBotResponse{
			Message: messagesOnThreadResponse.Data[0].Content[0].Text.Value,
			Role:    "assistant",
		}, nil
	} else {
		return nil, nil
	}
}
