package zerologlint

import (
	"github.com/ykadowak/zerologlint"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(zerologlint.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
