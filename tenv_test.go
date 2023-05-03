package tenv_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/sivchari/tenv"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, tenv.Analyzer, "a")
}

func TestAnalyzerGo116(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	a := tenv.Analyzer
	a.Flags.Parse([]string{"-go", "1.16"})
	analysistest.Run(t, testdata, a, "go116")
}
