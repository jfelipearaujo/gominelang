package patchouli_book

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_category"
	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_entry"
)

const (
	CATEGORIES_FOLDER string = "categories"
	ENTRIES_FOLDER    string = "entries"
)

type service struct {
	patchouliCategory patchouli_category.Service
	patchouliEntry    patchouli_entry.Service

	fromLang string
	toLang   string
}

func New(
	patchouliCategory patchouli_category.Service,
	patchouliEntry patchouli_entry.Service,
) Service {
	return &service{
		patchouliCategory: patchouliCategory,
		patchouliEntry:    patchouliEntry,
	}
}

func (s *service) SetLang(fromLang string, toLang string) {
	s.fromLang = fromLang
	s.toLang = toLang

	s.patchouliCategory.SetLang(fromLang, toLang)
	s.patchouliEntry.SetLang(fromLang, toLang)
}

func (s *service) Translate(inputFolder string, outputFolder string) error {
	inputFolder = fmt.Sprintf("%s/%s", inputFolder, s.fromLang)
	outputFolder = fmt.Sprintf("%s/%s", outputFolder, s.toLang)

	fmt.Printf("\nProcessing patchouli book...\n")

	err := filepath.WalkDir(inputFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk input folder '%s': %w", inputFolder, err)
		}

		if d.IsDir() {
			return nil
		}

		parentFolders := strings.Split(strings.TrimSuffix(path, filepath.Base(path)), "/")
		parentFolder := parentFolders[len(parentFolders)-2]

		output := fmt.Sprintf("%s/%s", outputFolder, parentFolder)

		if _, err := os.Stat(output); os.IsNotExist(err) {
			if err := os.MkdirAll(output, 0755); err != nil {
				return fmt.Errorf("failed to create output folder '%s': %w", output, err)
			}
		}

		outputFile := fmt.Sprintf("%s/%s", output, d.Name())

		switch strings.ToLower(parentFolder) {
		case CATEGORIES_FOLDER:
			if err := s.patchouliCategory.Translate(path, outputFile); err != nil {
				return err
			}
		case ENTRIES_FOLDER:
			if err := s.patchouliEntry.Translate(path, outputFile); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported folder '%s'", parentFolder)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to walk input folder '%s': %w", inputFolder, err)
	}

	return nil
}
