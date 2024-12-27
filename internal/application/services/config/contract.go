package config

import (
	"github.com/jfelipearaujo/gominelang/internal/application/services/translation/engine"
	"github.com/jfelipearaujo/gominelang/internal/domain"
)

// Service responsible to handle the config file
//
// Example:
//   - configPath: .gominelang.yaml
type Service interface {
	// Read and parse the config file and returns the domain.Config
	Read(configPath string) (*domain.Config, error)

	// Returns the translation engine that will be used
	GetEngine(config *domain.Config) (engine.Service, error)
}
