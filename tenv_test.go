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

func TestRun(t *testing.T) {
	t.Run("empty Go version", func(t *testing.T) {
		for _, goVersion := range []string{
			"", "go",
		} {
			testdata := testutil.WithModules(t, analysistest.TestData(), nil)
			a := tenv.Analyzer
			a.Flags.Parse([]string{"-go", goVersion})
			analysistest.Run(t, testdata, a, "a")
		}
	})

	t.Run("invalid Go version", func(t *testing.T) {
		for _, goVersion := range []string{
			"go1", "goa.2", "go1.a",
		} {
			a := tenv.Analyzer
			a.Flags.Parse([]string{"-go", goVersion})
			_, err := a.Run(nil)

			if err == nil {
				t.Error("expected error, but got <nil>")
			}
		}
	})
}
