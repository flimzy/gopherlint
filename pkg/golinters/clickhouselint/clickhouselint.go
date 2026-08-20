package clickhouselint

import (
	"github.com/ClickHouse/clickhouse-go-linter/passes/chbatchclose"
	"github.com/ClickHouse/clickhouse-go-linter/passes/chrowserr"
	"golang.org/x/tools/go/analysis"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"clickhouselint",
		"Detects common mistakes with the ClickHouse native Go driver API.",
		[]*analysis.Analyzer{chrowserr.NewAnalyzer(), chbatchclose.NewAnalyzer()},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
