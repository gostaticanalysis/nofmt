package nofmt_test

import (
	"testing"

	"github.com/gostaticanalysis/nofmt"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, nofmt.Analyzer, "a")
}
