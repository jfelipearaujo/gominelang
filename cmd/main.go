package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/jfelipearaujo/gominelang/internal/application/controllers/translator"
	"github.com/jfelipearaujo/gominelang/internal/application/services/config"
	"github.com/jfelipearaujo/gominelang/internal/application/services/db"
	"github.com/jfelipearaujo/gominelang/internal/application/services/lang"
	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_book"
	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_category"
	"github.com/jfelipearaujo/gominelang/internal/application/services/patchouli_entry"
	"github.com/jfelipearaujo/gominelang/internal/application/services/tag"
)

// GoMineLangVersion is the version of the CLI to be overwritten by build command
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

	config, err := configService.Read(CONFIG_FILE_NAME)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	dbService := db.New()
	if err := dbService.Open(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer dbService.Close()

	engine, err := configService.GetEngine(config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tagService := tag.New(engine)

	langService := lang.New(dbService, engine)

	patchouliCategoryService := patchouli_category.New(dbService, tagService)
	patchouliEntryService := patchouli_entry.New(dbService, engine, tagService)
	patchouliBookService := patchouli_book.New(patchouliCategoryService, patchouliEntryService)

	translatorController := translator.New(configService, langService, patchouliBookService)

	translatorController.Handle(context.Background(), config)
}
