package tag

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/jfelipearaujo/gominelang/internal/application/services/translation/engine"
)

var tagRegex = regexp.MustCompile(`(\$\(l:[a-zA-Z_]+:[a-zA-Z_]+\))`)

type service struct {
	translate engine.Service

	fromLang string
	toLang   string
}

func New(translate engine.Service) Service {
	return &service{
		translate: translate,
	}
}

func (s *service) SetLang(fromLang string, toLang string) {
	s.fromLang = fromLang
	s.toLang = toLang
}

func (s *service) HandleTranslation(input interface{}) error {
	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				if elem.Kind() == reflect.Ptr && !elem.IsNil() {
					elem = elem.Elem()
				}

				if elem.Kind() == reflect.Struct {
					if err := s.HandleTranslation(elem.Addr().Interface()); err != nil {
						return err
					}
				} else if elem.Kind() == reflect.String {
					tag := fieldType.Tag.Get("translate")
					if tag == "true" {
						translation, err := s.translate.Translate(s.fromLang, s.toLang, elem.String())
						if err != nil {
							return fmt.Errorf("failed to translate '%s' from '%s' to '%s': %w", elem.String(), s.fromLang, s.toLang, err)
						}
						elem.SetString(translation)
					}
				}
			}
			continue
		}

		tag := fieldType.Tag.Get("translate")
		if tag == "true" {
			if field.Kind() == reflect.String {
				translation, err := s.translate.Translate(s.fromLang, s.toLang, field.String())
				if err != nil {
					return fmt.Errorf("failed to translate '%s' from '%s' to '%s': %w", field.String(), s.fromLang, s.toLang, err)
				}
				field.SetString(translation)
			} else if !field.IsNil() {
				strVal := field.Elem()
				translation, err := s.translate.Translate(s.fromLang, s.toLang, strVal.String())
				if err != nil {
					return fmt.Errorf("failed to translate '%s' from '%s' to '%s': %w", strVal, s.fromLang, s.toLang, err)
				}
				strVal.SetString(translation)
			}
		}
	}

	return nil
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
