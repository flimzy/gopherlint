package gofmt

import (
	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/goformatters"
	gofmtbase "github.com/flimzy/gopherlint/pkg/goformatters/gofmt"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

func New(settings *config.GoFmtSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(gofmtbase.Name),
				"Check if the code is formatted according to 'gofmt' command.",
				gofmtbase.New(settings),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
