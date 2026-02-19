package nosprintfhostport

import (
	"github.com/stbenjam/no-sprintf-host-port/pkg/analyzer"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.Analyzer).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
