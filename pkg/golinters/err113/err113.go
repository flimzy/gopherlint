package err113

import (
	"github.com/Djarvur/go-err113"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(err113.NewAnalyzer()).
		WithDesc("Check errors handling expressions").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
