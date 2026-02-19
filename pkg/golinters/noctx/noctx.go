package noctx

import (
	"github.com/sonatard/noctx"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(noctx.Analyzer).
		WithDesc("Detects function and method with missing usage of context.Context").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
