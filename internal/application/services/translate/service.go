package translate

import (
	"regexp"
	"strings"

	gt "github.com/bas24/googletranslatefree"
)

var tagRegex = regexp.MustCompile(`(\$\(l:[a-zA-Z_]+:[a-zA-Z_]+\))`)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) Translate(from string, to string, text string) (string, error) {
	return gt.Translate(text, from, to)
}

func (s *service) FixWrongTranslation(original string, translated string) string {
	originalTags := []string{}
	matches := tagRegex.FindAllStringSubmatch(original, -1)
	for _, match := range matches {
		if len(match) > 1 {
			originalTags = append(originalTags, match[1])
		}
	}

	translatedTags := []string{}
	matches = tagRegex.FindAllStringSubmatch(translated, -1)
	for _, match := range matches {
		if len(match) > 1 {
			translatedTags = append(translatedTags, match[1])
		}
	}

	if len(originalTags) != len(translatedTags) {
		return translated
	}

	result := translated

	for i, originalTag := range originalTags {
		translatedTag := translatedTags[i]

		if originalTag != translatedTag {
			result = strings.ReplaceAll(result, translatedTag, originalTag)
		}
	}

	return result
}
