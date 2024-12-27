package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-version"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translation/engine"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translation/google_translate"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translation/open_ai"
	"github.com/jfelipearaujo/gominelang/internal/domain"
	"gopkg.in/yaml.v2"
)

const (
	// Update this version when you make a breaking change to the config file
	// and want to force users to update their config file
	MIN_SUPPORTED_VERSION string = "0.0.1"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) Read(configPath string) (*domain.Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := domain.Config{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(&config); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	configVersion, err := version.NewVersion(config.Version)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file version: %w", err)
	}

	minSupportedVersion, err := version.NewVersion(MIN_SUPPORTED_VERSION)
	if err != nil {
		return nil, fmt.Errorf("failed to parse min supported version: %w", err)
	}

	if configVersion.LessThan(minSupportedVersion) {
		return nil, fmt.Errorf("config file version is too old, please update to %s or higher", MIN_SUPPORTED_VERSION)
	}

	if err := s.checkEngine(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (s *service) checkEngine(config *domain.Config) error {
	if config.Engine.GoogleTranslate != nil && config.Engine.OpenAI != nil {
		if !config.Engine.GoogleTranslate.Enabled && !config.Engine.OpenAI.Enabled {
			return fmt.Errorf("at least one engine must be enabled")
		}

		if config.Engine.GoogleTranslate.Enabled && config.Engine.OpenAI.Enabled {
			return fmt.Errorf("only one engine can be enabled at a time")
		}
	}

	return nil
}

func (s *service) GetEngine(config *domain.Config) (engine.Service, error) {
	if config.Engine.GoogleTranslate != nil && config.Engine.GoogleTranslate.Enabled {
		return google_translate.New(), nil
	}

	if config.Engine.OpenAI != nil && config.Engine.OpenAI.Enabled {
		return open_ai.New(config), nil
	}

	return nil, fmt.Errorf("no engine is enabled")
}
