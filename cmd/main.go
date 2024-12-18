package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/jfelipearaujo/gominelang/internal/application/services/base_lang"
	"github.com/jfelipearaujo/gominelang/internal/application/services/config"
	"github.com/jfelipearaujo/gominelang/internal/application/services/dbhash"
	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_book"
	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_category"
	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_entry"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translate"
	"github.com/jfelipearaujo/gominelang/internal/application/services/translate_tag"
)

// GoMineLangVersion is the version of the CLI to be overwritten by goreleaser in the CI run with the version of the release in github
var GoMineLangVersion string

const (
	CONFIG_FILE_NAME string = ".gominelang.yaml"
)

func getVersion() string {
	noVersionAvailable := "No version info available for this build"

	if len(GoMineLangVersion) > 0 {
		return GoMineLangVersion
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return noVersionAvailable
	}

	if bi.Main.Version != "(devel)" {
		return bi.Main.Version
	}

	var vcsRevision string
	var vcsTime time.Time
	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		}
	}

	if vcsRevision != "" {
		return fmt.Sprintf("%s (%s)", vcsRevision, vcsTime)
	}

	return noVersionAvailable
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("Version: %s\n", getVersion())
		os.Exit(0)
	}

	configService := config.New()

	dbhashService := dbhash.New()
	if err := dbhashService.Open(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer dbhashService.Close()

	translateService := translate.New()
	translateTagService := translate_tag.New(translateService)

	baseLangsService := base_lang.New(dbhashService, translateService)

	patchouliCategoryService := patchouli_category.New(dbhashService, translateTagService)
	patchouliEntryService := patchouli_entry.New(dbhashService, translateService, translateTagService)
	patchouliBookService := patchouli_book.New(patchouliCategoryService, patchouliEntryService)

	config, err := configService.Read(CONFIG_FILE_NAME)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Version: %s\n", config.Version)

	for name, translate := range config.Translate {
		fmt.Printf("Mod: %s\n", name)
		fmt.Printf("Translating from '%s' to '%s'\n", translate.From, translate.To)

		baseLangsService.SetLang(translate.From, translate.To)

		if translate.Lang != nil {
			baseLangsService.SetLang(translate.From, translate.To)

			if err := baseLangsService.Translate(translate.Lang.Input, translate.Lang.Output); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}

		if translate.PatchouliBook != nil {
			patchouliBookService.SetLang(translate.From, translate.To)

			if err := patchouliBookService.Translate(translate.PatchouliBook.Input, translate.PatchouliBook.Output); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}

	fmt.Printf("\nDone!\n")
}
