package a

import (
	"os"
	"testing"
)

var (
	e   = os.Setenv("a", "b") // -all = true, want "variable e is not using testing.Setenv"
	_   = e
	env string
)

func setup() {
	os.Setenv("a", "b")        // -all = true, want "func setup is not using testing.Setenv"
	err := os.Setenv("a", "b") // -all = true, want "func setup is not using testing.Setenv"
	if err != nil {
		_ = err
	}
	env = os.Getenv("a")
	os.Setenv("a", "b") // -all = true, "func setup is not using testing.Setenv"
}

func TestF(t *testing.T) {
	setup()
	os.Setenv("a", "b")                         // want "func TestF is not using testing.Setenv"
	if err := os.Setenv("a", "b"); err != nil { // want "func TestF is not using testing.Setenv"
		_ = err
	}
}

func BenchmarkF(b *testing.B) {
	testTB(b)
	os.Setenv("a", "b")                         // want "func BenchmarkF is not using testing.Setenv"
	if err := os.Setenv("a", "b"); err != nil { // want "func BenchmarkF is not using testing.Setenv"
		_ = err
	}
}

func testTB(tb testing.TB) {
	os.Setenv("a", "b") // want "func testTB is not using testing.Setenv"
}
