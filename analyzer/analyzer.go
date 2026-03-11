package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc:  "linter for log messages",
	Run:  run,
}
