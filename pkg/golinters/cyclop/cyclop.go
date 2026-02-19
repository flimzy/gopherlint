package cyclop

import (
	"github.com/bkielbasa/cyclop/pkg/analyzer"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.CyclopSettings) *goanalysis.Linter {
	cfg := map[string]any{}

	if settings != nil {
		// Should be managed with `linters.exclusions.rules`.
		cfg["skipTests"] = false

		if settings.MaxComplexity != 0 {
			cfg["maxComplexity"] = settings.MaxComplexity
		}

		if settings.PackageAverage != 0 {
			cfg["packageAverage"] = settings.PackageAverage
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
