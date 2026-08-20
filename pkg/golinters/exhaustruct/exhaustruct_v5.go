package exhaustruct

import (
	exhaustruct "dev.gaijin.team/go/exhaustruct/v5/analyzer"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
	"github.com/flimzy/gopherlint/pkg/golinters/internal"
)

func NewV5(settings *config.ExhaustructV5Settings) *goanalysis.Linter {
	cfg := exhaustruct.Config{}

	if settings != nil {
		cfg.EnforcePatterns = settings.EnforcePatterns
		cfg.IgnorePatterns = settings.IgnorePatterns
		cfg.OptionalPatterns = settings.OptionalPatterns
		cfg.AllowEmpty = settings.AllowEmpty
		cfg.AllowEmptyPatterns = settings.AllowEmptyPatterns
		cfg.AllowEmptyReturns = settings.AllowEmptyReturns
		cfg.AllowEmptyDeclarations = settings.AllowEmptyDeclarations
		cfg.ExplicitMode = settings.ExplicitMode
	}

	analyzer, err := exhaustruct.NewAnalyzerWithConfig(cfg)
	if err != nil {
		internal.LinterLogger.Fatalf("exhaustruct configuration: %v", err)
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer).
		WithVersion(5). //nolint:mnd // It's the linter version.
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
