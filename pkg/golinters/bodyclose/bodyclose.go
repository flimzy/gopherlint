package bodyclose

import (
	"github.com/timakin/bodyclose/passes/bodyclose"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(bodyclose.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
