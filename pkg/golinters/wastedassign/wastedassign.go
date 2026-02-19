package wastedassign

import (
	"github.com/sanposhiho/wastedassign/v2"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(wastedassign.Analyzer).
		WithDesc("Finds wasted assignment statements").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
