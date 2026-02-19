package testableexamples

import (
	"github.com/maratori/testableexamples/pkg/testableexamples"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(testableexamples.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
