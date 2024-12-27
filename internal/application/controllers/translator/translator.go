package translator

import (
	"context"
	"fmt"
	"os"

	"github.com/jfelipearaujo/gominelang/internal/application/services/config"
	"github.com/jfelipearaujo/gominelang/internal/application/services/lang"
	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_book"
	"github.com/jfelipearaujo/gominelang/internal/domain"
)

type Controller struct {
	configService config.Service

	langService          lang.Service
	patchouliBookService patchouli_book.Service
}

func New(
	configService config.Service,
	langService lang.Service,
	patchouliBookService patchouli_book.Service,
) *Controller {
	return &Controller{
		configService:        configService,
		langService:          langService,
		patchouliBookService: patchouliBookService,
	}
}

func (ct *Controller) Handle(ctx context.Context, config *domain.Config) {
	fmt.Printf("Version: %s\n", config.Version)

	for name, translate := range config.Translate {
		fmt.Printf("Mod: %s\n", name)
		fmt.Printf("Translating from '%s' to '%s'\n", translate.From, translate.To)

		ct.langService.SetLang(translate.From, translate.To)

		if translate.Lang != nil {
			ct.langService.SetLang(translate.From, translate.To)

			if err := ct.langService.Translate(translate.Lang.Input, translate.Lang.Output); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}

		if translate.PatchouliBook != nil {
			ct.patchouliBookService.SetLang(translate.From, translate.To)

			if err := ct.patchouliBookService.Translate(translate.PatchouliBook.Input, translate.PatchouliBook.Output); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}

	fmt.Printf("\nDone!\n")
}
