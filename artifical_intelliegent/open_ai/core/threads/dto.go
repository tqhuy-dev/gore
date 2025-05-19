package threads

import (
	"github.com/tqhuy-dev/gore/artifical_intelliegent/open_ai/core/messages"
)

type CreateThreadDto struct {
	Messages []messages.MessageDto `json:"messages"`
}

type CreateThreadResponseDto struct {
	Id string `json:"id"`
}

type RunThreadRequest struct {
	ThreadId    string `json:"thread_id"`
	AssistantId string `json:"assistant_id"`
}

type RetrieveThreadRunRequest struct {
	ThreadId string `json:"thread_id"`
	RunId    string `json:"run_id"`
}

type RetrieveThreadRunResponse struct {
	ThreadId    string `json:"thread_id"`
	AssistantId string `json:"assistant_id"`
	Status      string `json:"status"`
}

type SendMessageThreadsRequest struct {
	ThreadId string `json:"thread_id"`
	Content  string `json:"content"`
	Role     string `json:"role"`
}

type SendMessageThreadsResponse struct {
	ThreadId string `json:"thread_id"`
	Id       string `json:"id"`
}

type ListMessagesOnThreadRequest struct {
	ThreadId string `json:"thread_id"`
}

type ListMessagesOnThreadResponse struct {
	Data []ListMessagesDataOnThreadResponse `json:"data"`
}

type MessageContentThreadsData struct {
	Type string `json:"type"`
	Text struct {
		Value string `json:"value"`
	} `json:"text"`
}
type ListMessagesDataOnThreadResponse struct {
	Id      string                      `json:"id"`
	Content []MessageContentThreadsData `json:"content"`
}
