package assistants

type CreateAssistantRequest struct {
	Instruction string   `json:"instruction"`
	Name        string   `json:"name"`
	Tools       []string `json:"tools"`
	Model       string   `json:"model"`
}

type CreateAssistantResponse struct {
	Id string `json:"id"`
}

type ModifyAssistantRequest struct {
	AssistantId string `json:"assistant_id"`
}

type ModifyAssistantResponse struct {
	Id string `json:"id"`
}

type DeleteAssistantRequest struct {
	AssistantId string `json:"assistant_id"`
}

type DeleteAssistantResponse struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

type ListAssistantRequest struct {
	Limit int `json:"limit"`
}

type ListAssistantData struct {
	Id           string   `json:"id"`
	Instructions string   `json:"instructions"`
	Tools        []string `json:"tools"`
	Model        string   `json:"model"`
}
type ListAssistantResponse struct {
	Data []ListAssistantData `json:"data"`
}
