package interfacebloat

import (
	"github.com/sashamelentyev/interfacebloat/pkg/analyzer"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.InterfaceBloatSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			analyzer.InterfaceMaxMethodsFlag: settings.Max,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.New()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
