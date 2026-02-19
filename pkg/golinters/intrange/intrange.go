package intrange

import (
	"github.com/ckaznocha/intrange"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(intrange.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
