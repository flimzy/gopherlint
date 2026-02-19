package golines

import (
	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/goformatters"
	golinesbase "github.com/flimzy/gopherlint/pkg/goformatters/golines"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

func New(settings *config.GoLinesSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(golinesbase.Name),
				"Checks if code is formatted, and fixes long lines",
				golinesbase.New(settings),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
