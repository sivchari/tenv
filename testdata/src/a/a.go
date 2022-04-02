package a

import (
	"os"
	"testing"
)

var (
	ee = os.Setenv("a", "b") // never seen
)

func setup() {
	os.Setenv("a", "b")        // never seen
	err := os.Setenv("a", "b") // never seen
	if err != nil {
		_ = err
	}
	os.Setenv("a", "b") // never seen
}

func F(t *testing.T) {
	setup()
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in F"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in F"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `t\\.Setenv\\(\\)` in F"
		_ = err
	}
}

func BF(b *testing.B) {
	TBF(b)
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BF"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BF"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `b\\.Setenv\\(\\)` in BF"
		_ = err
	}
}

func TBF(tb testing.TB) {
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TBF"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TBF"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `tb\\.Setenv\\(\\)` in TBF"
		_ = err
	}
}

func FF(f *testing.F) {
	os.Setenv("a", "b")        // want "os\\.Setenv\\(\\) can be replaced by `f\\.Setenv\\(\\)` in FF"
	err := os.Setenv("a", "b") // want "os\\.Setenv\\(\\) can be replaced by `f\\.Setenv\\(\\)` in FF"
	_ = err
	if err := os.Setenv("a", "b"); err != nil { // want "os\\.Setenv\\(\\) can be replaced by `f\\.Setenv\\(\\)` in FF"
		_ = err
	}
}
