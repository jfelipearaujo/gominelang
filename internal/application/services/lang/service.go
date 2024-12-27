package lang

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jfelipearaujo/gominelang/internal/application/services/db"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translation/engine"
	"github.com/jfelipearaujo/gominelang/internal/domain"
)

type service struct {
	db        db.Service
	translate engine.Service

	fromLang string
	toLang   string
}

func New(db db.Service, translate engine.Service) Service {
	return &service{
		db:        db,
		translate: translate,
	}
}

func (s *service) SetLang(fromLang string, toLang string) {
	s.fromLang = fromLang
	s.toLang = toLang
}

func (s *service) Translate(inputFolder string, outputFolder string) error {
	inputFile := fmt.Sprintf("%s/%s.json", inputFolder, s.fromLang)
	outputFile := fmt.Sprintf("%s/%s.json", outputFolder, s.toLang)

	inputItems := make(domain.GameItems)
	outputItems := make(domain.GameItems)

	fmt.Printf("\nProcessing base lang file...\n")

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file '%s' not found", inputFile)
	}

	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file '%s': %w", inputFile, err)
	}

	if err := json.Unmarshal(inputData, &inputItems); err != nil {
		return fmt.Errorf("failed to unmarshal input file '%s': %w", inputFile, err)
	}

	hashExists, err := s.db.Exists(inputFile)
	if err != nil {
		return fmt.Errorf("failed to check if the file '%s' exists in the database: %w", inputFile, err)
	}

	if hashExists != nil {
		equals, err := s.db.Compare(hashExists, inputFile)
		if err != nil {
			return fmt.Errorf("failed to compare the hash of the file '%s' with the database: %w", inputFile, err)
		}

		if equals {
			fmt.Printf("Skipping base lang file, no changes detected\n")
			return nil
		}
	}

	fmt.Printf("Changes detected, translating...\n")

	if _, err := os.Stat(outputFile); err == nil {
		outputData, err := os.ReadFile(outputFile)
		if err != nil {
			return fmt.Errorf("failed to read output file '%s': %w", outputFile, err)
		}

		if err := json.Unmarshal(outputData, &outputItems); err != nil {
			return fmt.Errorf("failed to unmarshal output file '%s': %w", outputFile, err)
		}
	}

	for key, value := range inputItems {
		if _, exists := outputItems[key]; !exists {
			translation, err := s.translate.Translate(s.fromLang, s.toLang, value)
			if err != nil {
				return fmt.Errorf("failed to translate '%s' from '%s' to '%s': %w", value, s.fromLang, s.toLang, err)
			}

			outputItems[key] = translation
		}
	}

	outputData, err := json.MarshalIndent(outputItems, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal output file '%s': %w", outputFile, err)
	}

	if err := os.WriteFile(outputFile, outputData, 0644); err != nil {
		return fmt.Errorf("failed to write output file '%s': %w", outputFile, err)
	}

	if err := s.db.Store(inputFile); err != nil {
		return fmt.Errorf("failed to store the hash of the file '%s' in the database: %w", inputFile, err)
	}

	return nil
}
