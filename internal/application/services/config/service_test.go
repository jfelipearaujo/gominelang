package config_test

import (
	"testing"

	"github.com/jfelipearaujo/gominelang/internal/application/services/config"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	t.Run("Should read config file successfully", func(t *testing.T) {
		// Arrange
		configFiles := []string{
			"./testdata/full_config.yaml",
			"./testdata/only_lang.yaml",
			"./testdata/only_patchouli.yaml",
		}

		sut := config.New()

		for _, configFile := range configFiles {
			// Act
			config, err := sut.Read(configFile)

			// Assert
			assert.NoError(t, err)
			assert.NotNil(t, config)
		}
	})

	t.Run("Should return validation error when read invalid config file", func(t *testing.T) {
		// Arrange
		configFile := "./testdata/invalid_lang.yaml"

		sut := config.New()

		// Act
		config, err := sut.Read(configFile)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, config)
	})

	t.Run("Should return error when config file is not found", func(t *testing.T) {
		// Arrange
		configFile := "./testdata/not_found.yaml"

		sut := config.New()

		// Act
		config, err := sut.Read(configFile)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, config)
	})

	t.Run("Should return error when read malformed config", func(t *testing.T) {
		// Arrange
		configFile := "./testdata/malformed.yaml"

		sut := config.New()

		// Act
		config, err := sut.Read(configFile)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, config)
	})

	t.Run("Should return error when read old config file", func(t *testing.T) {
		// Arrange
		configFile := "./testdata/old_config.yaml"

		sut := config.New()

		// Act
		config, err := sut.Read(configFile)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, config)
	})
}
