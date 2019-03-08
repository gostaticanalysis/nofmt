package main

import (
	"github.com/gostaticanalysis/nofmt"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(nofmt.Analyzer) }
