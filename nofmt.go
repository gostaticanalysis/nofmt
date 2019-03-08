package nofmt

import (
	"go/format"
	"io/ioutil"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "nofmt",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "nofmt checks neither go files formated or not"

func run(pass *analysis.Pass) (interface{}, error) {
	for i := range pass.Files {
		pos := pass.Files[i].Pos()
		fname := pass.Fset.File(pos).Name()
		src, err := ioutil.ReadFile(fname)
		if err != nil {
			return nil, err
		}
		dst, err := format.Source(src)
		if err != nil {
			return nil, err
		}
		if string(src) != string(dst) {
			pass.Reportf(pos, "%s is not formated by gofmt", fname)
		}
	}
	return nil, nil
}
