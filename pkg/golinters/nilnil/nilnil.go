package nilnil

import (
	"github.com/Antonboom/nilnil/pkg/analyzer"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.NilNilSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"detect-opposite": settings.DetectOpposite,
		}
		if b := settings.OnlyTwo; b != nil {
			cfg["only-two"] = *b
		}
		if len(settings.CheckedTypes) != 0 {
			cfg["checked-types"] = settings.CheckedTypes
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.New()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
