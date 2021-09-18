package a

import (
	"os"
	"testing"
)

var (
	e = os.Setenv("a", "b") // never seen
)

func setup() {
	os.Setenv("a", "b")        // if -all = true, want "func setup is not using testing.Setenv"
	err := os.Setenv("a", "b") // if -all = true, want "func setup is not using testing.Setenv"
	if err != nil {
		_ = err
	}
	os.Setenv("a", "b") // if -all = true, "func setup is not using testing.Setenv"
}

func TestF(t *testing.T) {
	setup()
	os.Setenv("a", "b")                         // want "os.Setenv() can be replaced by `t.Setenv()` in TestF"
	if err := os.Setenv("a", "b"); err != nil { // want "os.Setenv() can be replaced by `t.Setenv()` in TestF"
		_ = err
	}
}

func BenchmarkF(b *testing.B) {
	testTB(b)
	os.Setenv("a", "b")                         // want "os.Setenv() can be replaced by `b.Setenv()` in BenchmarkF"
	if err := os.Setenv("a", "b"); err != nil { // want "os.Setenv() can be replaced by `b.Setenv()` in BenchmarkF"
		_ = err
	}
}

func testTB(tb testing.TB) {
	os.Setenv("a", "b") // want "os.Setenv() can be replaced by `tb.Setenv()` in testTB"
}
