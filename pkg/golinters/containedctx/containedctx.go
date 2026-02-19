package containedctx

import (
	"github.com/sivchari/containedctx"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(containedctx.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
