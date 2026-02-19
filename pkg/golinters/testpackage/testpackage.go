package testpackage

import (
	"strings"

	"github.com/maratori/testpackage/pkg/testpackage"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.TestpackageSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			testpackage.SkipRegexpFlagName:    settings.SkipRegexp,
			testpackage.AllowPackagesFlagName: strings.Join(settings.AllowPackages, ","),
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(testpackage.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
