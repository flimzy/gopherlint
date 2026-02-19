package asasalint

import (
	"github.com/alingse/asasalint"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

func New(settings *config.AsasalintSettings) *goanalysis.Linter {
	cfg := asasalint.LinterSetting{}
	if settings != nil {
		cfg.Exclude = settings.Exclude
		cfg.NoBuiltinExclusions = !settings.UseBuiltinExclusions

		// Should be managed with `linters.exclusions.rules`.
		cfg.IgnoreTest = false
	}

	analyzer, err := asasalint.NewAnalyzer(cfg)
	if err != nil {
		internal.LinterLogger.Fatalf("asasalint: create analyzer: %v", err)
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
