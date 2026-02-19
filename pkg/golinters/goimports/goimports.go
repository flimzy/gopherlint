package goimports

import (
	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/goformatters"
	goimportsbase "github.com/flimzy/gopherlint/pkg/goformatters/goimports"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

func New(settings *config.GoImportsSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(goimportsbase.Name),
				"Checks if the code and import statements are formatted according to the 'goimports' command.",
				goimportsbase.New(settings),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
