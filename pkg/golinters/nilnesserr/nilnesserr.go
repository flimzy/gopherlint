package nilnesserr

import (
	"github.com/alingse/nilnesserr"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

func New() *goanalysis.Linter {
	analyzer, err := nilnesserr.NewAnalyzer(nilnesserr.LinterSetting{})
	if err != nil {
		internal.LinterLogger.Fatalf("nilnesserr: create analyzer: %v", err)
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
