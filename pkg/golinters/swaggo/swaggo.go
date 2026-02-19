package swaggo

import (
	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/goformatters"
	"github.com/flimzy/gopherlint/pkg/goformatters/swaggo"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(swaggo.Name),
				"Check if swaggo comments are formatted",
				swaggo.New(),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
