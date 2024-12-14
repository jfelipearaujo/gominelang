package config

import (
	"github.com/jfelipearaujo/gominelang/internal/domain"
)

// Service responsible to handle the config file
//
// Example:
//   - configPath: .gominelang.yaml
type Service interface {
	// Read and parse the config file and returns the domain.Config
	Read(configPath string) (*domain.Config, error)
}
