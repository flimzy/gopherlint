package gofumpt

import (
	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/goformatters"
	gofumptbase "github.com/flimzy/gopherlint/pkg/goformatters/gofumpt"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

func New(settings *config.GoFumptSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(gofumptbase.Name),
				"Check if code and import statements are formatted, with additional rules.",
				gofumptbase.New(settings, settings.LangVersion),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
