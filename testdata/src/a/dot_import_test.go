package a

import (
	"fmt"
	. "os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	e3 = Setenv("a", "b") // never seen
)

func testsetup3() {
	Setenv("a", "b")        // if -all = true, want  "os\\.Setenv\\(\\) can be replaced by `testing\\.Setenv\\(\\)` in testsetup3"
	err := Setenv("a", "b") // if -all = true, want  "os\\.Setenv\\(\\) can be replaced by `testing\\.Setenv\\(\\)` in testsetup3"
	if err != nil {
		_ = err
	}
	Setenv("a", "b") // if -all = true, "func setup is not using testing.Setenv"
}

func TestF3(t *testing.T) {
	testsetup3()
	Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF3"
	err := Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF3"
	_ = err
	if err := Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestF3"
		_ = err
	}
}

func BenchmarkF3(b *testing.B) {
	TB(b)
	Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF3"
	err := Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF3"
	_ = err
	if err := Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BenchmarkF3"
		_ = err
	}
}

func TB3(tb testing.TB) {
	Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB3"
	err := Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB3"
	_ = err
	if err := Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TB3"
		_ = err
	}
}

func TestFunctionLiteral3(t *testing.T) {
	testsetup3()
	t.Run("test", func(t *testing.T) {
		Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
		err := Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
		_ = err
		if err := Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
			_ = err
		}
	})
}

func TestEmpty3(t *testing.T) {
	t.Run("test", func(*testing.T) {})
}

func TestEmptyTB3(t *testing.T) {
	func(testing.TB) {}(t)
}

func TestTDD3(t *testing.T) {
	for _, tt := range []struct {
		name string
	}{
		{"test"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
			err := Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
			_ = err
			if err := Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in anonymous function"
				_ = err
			}
		})
	}
}

func TestLoop3(t *testing.T) {
	for i := 0; i < 3; i++ {
		Setenv(fmt.Sprintf("a%d", i), "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop3"
		err := Setenv(fmt.Sprintf("a%d", i), "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop3"
		_ = err
		if err := Setenv(fmt.Sprintf("a%d", i), "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestLoop3"
			_ = err
		}
	}
}

func TestUsingArg3(t *testing.T) {
	require.NoError(t, Setenv("a", "b")) // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in TestUsingArg3"
}
