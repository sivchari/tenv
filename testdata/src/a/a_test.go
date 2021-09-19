package a

import (
	"os"
	"testing"
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
