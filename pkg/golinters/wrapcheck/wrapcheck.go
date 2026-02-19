package wrapcheck

import (
	"github.com/tomarrell/wrapcheck/v2/wrapcheck"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.WrapcheckSettings) *goanalysis.Linter {
	cfg := wrapcheck.NewDefaultConfig()

	if settings != nil {
		cfg.ExtraIgnoreSigs = settings.ExtraIgnoreSigs
		cfg.ReportInternalErrors = settings.ReportInternalErrors

		if len(settings.IgnoreSigs) != 0 {
			cfg.IgnoreSigs = settings.IgnoreSigs
		}
		if len(settings.IgnoreSigRegexps) != 0 {
			cfg.IgnoreSigRegexps = settings.IgnoreSigRegexps
		}
		if len(settings.IgnorePackageGlobs) != 0 {
			cfg.IgnorePackageGlobs = settings.IgnorePackageGlobs
		}
		if len(settings.IgnoreInterfaceRegexps) != 0 {
			cfg.IgnoreInterfaceRegexps = settings.IgnoreInterfaceRegexps
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(wrapcheck.NewAnalyzer(cfg)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
