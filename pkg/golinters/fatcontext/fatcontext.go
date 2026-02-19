package fatcontext

import (
	"go.augendre.info/fatcontext/pkg/analyzer"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.FatcontextSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			analyzer.FlagCheckStructPointers: settings.CheckStructPointers,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
