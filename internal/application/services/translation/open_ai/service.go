package open_ai

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translation/engine"
	"github.com/jfelipearaujo/gominelang/internal/domain"
)

type service struct {
	config *domain.Config

	client *resty.Client
}

func New(config *domain.Config) engine.Service {
	return &service{
		config: config,
		client: resty.New().
			SetBaseURL("https://api.openai.com"),
	}
}

func (s *service) Translate(from string, to string, text string) (string, error) {
	request := NewRequest(from, to, text)

	var response Response

	result, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", s.config.Engine.OpenAI.APIKey)).
		SetBody(request).
		SetResult(&response).
		SetError(&response).
		Post("/v1/chat/completions")

	if err != nil {
		return "", fmt.Errorf("error to translate: %w", err)
	}

	if !result.IsSuccess() {
		return "", fmt.Errorf("failed to translate: %v", response.Error.Message)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no choices found")
	}

	if response.Choices[0].Message.Content == "" {
		return "", fmt.Errorf("no content found")
	}

	if response.Choices[0].Message.Role != ROLE_ASSISTANT {
		return "", fmt.Errorf("invalid role found")
	}

	translatedText := TranslatedText{}
	if err := json.Unmarshal([]byte(response.Choices[0].Message.Content), &translatedText); err != nil {
		return "", fmt.Errorf("failed to unmarshal translated text: %w", err)
	}

	return translatedText.Data, nil
}
