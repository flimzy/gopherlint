package gocheckcompilerdirectives

import (
	"4d63.com/gocheckcompilerdirectives/checkcompilerdirectives"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(checkcompilerdirectives.Analyzer()).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
