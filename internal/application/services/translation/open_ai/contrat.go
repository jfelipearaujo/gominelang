package open_ai

import "fmt"

const (
	ROLE_SYSTEM    string = "system"
	ROLE_USER      string = "user"
	ROLE_ASSISTANT string = "assistant"
)

type Request struct {
	Model          string           `json:"model"`
	Messages       []RequestMessage `json:"messages"`
	ResponseFormat ResponseFormat   `json:"response_format"`
}

type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseFormat struct {
	Type       string     `json:"type"`
	JsonSchema JsonSchema `json:"json_schema"`
}

type JsonSchema struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
}

type Schema struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
}

type Properties struct {
	TranslatedText TranslatedTextProperty `json:"translated_text"`
}

type TranslatedTextProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

func NewRequest(from string, to string, text string) Request {
	return Request{
		Model: "gpt-4o-2024-08-06",
		Messages: []RequestMessage{
			{
				Role:    ROLE_SYSTEM,
				Content: fmt.Sprintf("You are a experienced translator that can handle any kind of translation, now you will be responsible to translate texts from from '%s' to '%s'. The result must be returned into JSON format.", from, to),
			},
			{
				Role:    ROLE_USER,
				Content: fmt.Sprintf("Translate the following text from '%s' to '%s': %s", from, to, text),
			},
		},
		ResponseFormat: ResponseFormat{
			Type: "json_schema",
			JsonSchema: JsonSchema{
				Name: "translated_text_schema",
				Schema: Schema{
					Type: "object",
					Properties: Properties{
						TranslatedText: TranslatedTextProperty{
							Type:        "string",
							Description: "The translated text in the desired language",
						},
					},
					AdditionalProperties: false,
				},
			},
		},
	}
}

type Response struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`

	Error Error `json:"error"`
}

type TranslatedText struct {
	Data string `json:"translated_text"`
}

type Choice struct {
	Index        int             `json:"index"`
	Message      ResponseMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens        int `json:"prompt_tokens"`
	CompletionTokens    int `json:"completion_tokens"`
	TotalTokens         int `json:"total_tokens"`
	PromptTokensDetails struct {
		CachedTokens int `json:"cached_tokens"`
		AudioTokens  int `json:"audio_tokens"`
	} `json:"prompt_tokens_details"`
	CompletionTokensDetails struct {
		ReasoningTokens int `json:"reasoning_tokens"`
		AudioTokens     int `json:"audio_tokens"`
		AcceptedTokens  int `json:"accepted_tokens"`
		RejectedTokens  int `json:"rejected_tokens"`
	} `json:"completion_tokens_details"`
}

type Error struct {
	Code    string `json:"code"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Message string `json:"message"`
}
