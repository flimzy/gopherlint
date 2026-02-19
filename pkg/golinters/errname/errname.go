package errname

import (
	"github.com/Antonboom/errname/pkg/analyzer"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.New()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
