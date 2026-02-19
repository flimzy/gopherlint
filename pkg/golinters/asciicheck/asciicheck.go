package asciicheck

import (
	"github.com/golangci/asciicheck"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(asciicheck.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
