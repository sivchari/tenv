package a

import (
	"fmt"
	"os"
	"testing"
	mytest "testing"

	"github.com/stretchr/testify/require"
)

var (
	e2 = os.Setenv("a", "b") // never seen
)

func testsetup2() {
	os.Setenv("a", "b")        // if -all = true, want  "os\\.Setenv\\(\\) can be replaced by `testing\\.Setenv\\(\\)` in testsetup2"
	err := os.Setenv("a", "b") // if -all = true, want  "os\\.Setenv\\(\\) can be replaced by `testing\\.Setenv\\(\\)` in testsetup2"
	if err != nil {
		_ = err
	}
	os.Setenv("a", "b") // if -all = true, "func setup is not using testing.Setenv"
}

func TestF2(t *mytest.T) {
	testsetup2()
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF2"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF2"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF2"
		_ = err
	}
}

func BenchmarkF2(b *mytest.B) {
	TB(b)
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF2"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF2"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF2"
		_ = err
	}
}

func TB2(tb mytest.TB) {
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB2"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB2"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB2"
		_ = err
	}
}

func TestFunctionLiteral2(t *mytest.T) {
	testsetup()
	t.Run("test", func(t *mytest.T) {
		os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
		err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
		_ = err
		if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
			_ = err
		}
	})
}

func TestEmpty2(t *mytest.T) {
	t.Run("test", func(*testing.T) {})
}

func TestEmptyTB2(t *mytest.T) {
	func(testing.TB) {}(t)
}

func TestTDD2(t *mytest.T) {
	for _, tt := range []struct {
		name string
	}{
		{"test"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
			err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
			_ = err
			if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
				_ = err
			}
		})
	}
}

func TestLoop2(t *mytest.T) {
	for i := 0; i < 3; i++ {
		os.Setenv(fmt.Sprintf("a%d", i), "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop2"
		err := os.Setenv(fmt.Sprintf("a%d", i), "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop2"
		_ = err
		if err := os.Setenv(fmt.Sprintf("a%d", i), "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop2"
			_ = err
		}
	}
}

func TestUsingArg2(t *mytest.T) {
	require.NoError(t, os.Setenv("a", "b")) // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestUsingArg2"
}
