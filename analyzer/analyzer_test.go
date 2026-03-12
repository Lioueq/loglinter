package analyzer_test

import (
	"testing"

	"github.com/lioueq/loglinter/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analyzer.SetSensitivePatterns([]string{"password", "apikey", "api_key", "token", "creds"})
	t.Cleanup(func() {
		analyzer.SetSensitivePatterns(nil)
	})

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "basic")
}
