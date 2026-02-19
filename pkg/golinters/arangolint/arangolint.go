package arangolint

import (
	"go.augendre.info/arangolint/pkg/analyzer"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
