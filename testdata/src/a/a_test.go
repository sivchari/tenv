package a

import (
	"os"
	"testing"
)

var (
	e   = os.Setenv("a", "b") // want "variable e is not using t.Setenv"
	_   = e
	env string
)

func setup() {
	os.Setenv("a", "b")        // want "func setup is not using t.Setenv"
	err := os.Setenv("a", "b") // want "func setup is not using t.Setenv"
	if err != nil {
		_ = err
	}
	env = os.Getenv("a")
	os.Setenv("a", "b") // want "func setup is not using t.Setenv"
}

func TestF(t *testing.T) {
	setup()
	os.Setenv("a", "b")                         // want "func TestF is not using t.Setenv"
	if err := os.Setenv("a", "b"); err != nil { // want "func TestF is not using t.Setenv"
		_ = err
	}
}
