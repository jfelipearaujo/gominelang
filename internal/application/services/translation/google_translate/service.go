package google_translate

import (
	gt "github.com/bas24/googletranslatefree"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translation/engine"
)

type service struct{}

func New() engine.Service {
	return &service{}
}

func (s *service) Translate(from string, to string, text string) (string, error) {
	return gt.Translate(text, from, to)
}
