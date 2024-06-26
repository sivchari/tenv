package a

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	e = os.Setenv("a", "b") // never seen
)

func testsetup() {
	os.Setenv("a", "b")        // if -all = true, want  "os\\.Setenv\\(\\) can be replaced by `testing\\.Setenv\\(\\)` in testsetup"
	err := os.Setenv("a", "b") // if -all = true, want  "os\\.Setenv\\(\\) can be replaced by `testing\\.Setenv\\(\\)` in testsetup"
	if err != nil {
		_ = err
	}
	os.Setenv("a", "b") // if -all = true, "func setup is not using testing.Setenv"
}

func TestF(t *testing.T) {
	testsetup()
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF"
		_ = err
	}
}

func BenchmarkF(b *testing.B) {
	TB(b)
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF"
		_ = err
	}
}

func TB(tb testing.TB) {
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB"
		_ = err
	}
}

func TestFunctionLiteral(t *testing.T) {
	testsetup()
	t.Run("test", func(t *testing.T) {
		os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
		err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
		_ = err
		if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
			_ = err
		}
	})
}

func TestEmpty(t *testing.T) {
	t.Run("test", func(*testing.T) {})
}

func TestEmptyTB(t *testing.T) {
	func(testing.TB) {}(t)
}

func TestTDD(t *testing.T) {
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

func TestLoop(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.Setenv(fmt.Sprintf("a%d", i), "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop"
		err := os.Setenv(fmt.Sprintf("a%d", i), "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop"
		_ = err
		if err := os.Setenv(fmt.Sprintf("a%d", i), "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop"
			_ = err
		}
	}
}

func TestUsingArg(t *testing.T) {
	require.NoError(t, os.Setenv("a", "b")) // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestUsingArg"
}
