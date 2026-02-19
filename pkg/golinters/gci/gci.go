package gci

import (
	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/goformatters"
	gcibase "github.com/flimzy/gopherlint/pkg/goformatters/gci"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

const linterName = "gci"

func New(settings *config.GciSettings) *goanalysis.Linter {
	formatter, err := gcibase.New(settings)
	if err != nil {
		internal.LinterLogger.Fatalf("%s: create analyzer: %v", linterName, err)
	}

	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(linterName),
				"Check if code and import statements are formatted, with additional rules.",
				formatter,
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
