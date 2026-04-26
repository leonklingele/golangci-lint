package errgroupctx

import (
	errgroupctx "github.com/m-ocean-it/errgroup-ctx-lint"

	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(errgroupctx.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
