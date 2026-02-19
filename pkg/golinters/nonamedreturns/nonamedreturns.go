package nonamedreturns

import (
	"github.com/firefart/nonamedreturns/analyzer"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.NoNamedReturnsSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			analyzer.FlagReportErrorInDefer: settings.ReportErrorInDefer,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.Analyzer).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
