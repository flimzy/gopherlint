package sqlclosecheck

import (
	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewDeferOnlyAnalyzer()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
