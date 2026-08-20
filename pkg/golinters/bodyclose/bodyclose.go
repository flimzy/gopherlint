package bodyclose

import (
	"github.com/timakin/bodyclose/passes/bodyclose"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.BodyCloseSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"check-consumption": settings.CheckConsumption,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(bodyclose.Analyzer).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
