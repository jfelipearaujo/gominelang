package patchouli_entry

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/jfelipearaujo/gominelang/internal/application/services/dbhash"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translate"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translate_tag"
	"github.com/jfelipearaujo/gominelang/internal/domain"
)

type service struct {
	dbhash        dbhash.Service
	translate     translate.Service
	translate_tag translate_tag.Service
}

func New(
	dbhash dbhash.Service,
	translate translate.Service,
	translate_tag translate_tag.Service,
) Service {
	return &service{
		dbhash:        dbhash,
		translate:     translate,
		translate_tag: translate_tag,
	}
}

func (s *service) SetLang(fromLang string, toLang string) {
	s.translate_tag.SetLang(fromLang, toLang)
}

func (s *service) Translate(inputFile string, outputFile string) error {
	input := domain.PatchouliEntry{}
	output := domain.PatchouliEntry{}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file '%s' not found", inputFile)
	}

	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file '%s': %w", inputFile, err)
	}

	if err := json.Unmarshal(inputData, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input file '%s': %w", inputFile, err)
	}

	hashExists, err := s.dbhash.Exists(inputFile)
	if err != nil {
		return fmt.Errorf("failed to check if the file '%s' exists in the database: %w", inputFile, err)
	}

	if hashExists != nil {
		equals, err := s.dbhash.Compare(hashExists, inputFile)
		if err != nil {
			return fmt.Errorf("failed to compare the hash of the file '%s' with the database: %w", inputFile, err)
		}

		if equals {
			fmt.Printf("Skipping patchouli entry '%s', no changes detected\n", filepath.Base(inputFile))
			return nil
		}
	}

	fmt.Printf("Changes detected, translating...\n")

	if _, err := os.Stat(outputFile); err == nil {
		outputData, err := os.ReadFile(outputFile)
		if err != nil {
			return fmt.Errorf("failed to read output file '%s': %w", outputFile, err)
		}

		if err := json.Unmarshal(outputData, &output); err != nil {
			return fmt.Errorf("failed to unmarshal output file '%s': %w", outputFile, err)
		}
	}

	if reflect.DeepEqual(input, output) {
		return nil
	}

	output.MapFrom(&input)

	if err := s.translate_tag.HandleTranslation(&output); err != nil {
		return fmt.Errorf("failed to translate: %w", err)
	}

	for i := 0; i < len(output.Pages); i++ {
		if output.Pages[i].Title != nil {
			out := s.translate.FixWrongTranslation(*input.Pages[i].Title, *output.Pages[i].Title)
			output.Pages[i].Title = &out
		}

		if output.Pages[i].Text != nil {
			out := s.translate.FixWrongTranslation(*input.Pages[i].Text, *output.Pages[i].Text)
			output.Pages[i].Text = &out
		}
	}

	outputData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal output file '%s': %w", outputFile, err)
	}

	if err := os.WriteFile(outputFile, outputData, 0644); err != nil {
		return fmt.Errorf("failed to write output file '%s': %w", outputFile, err)
	}

	if err := s.dbhash.Store(inputFile); err != nil {
		return fmt.Errorf("failed to store the hash of the file '%s' in the database: %w", inputFile, err)
	}

	return nil
}
