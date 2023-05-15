package go116

import (
	"os"
	"testing"
)

func TestF(t *testing.T) {
	os.Setenv("a", "b")        // if -go = 1.16, ""
	err := os.Setenv("a", "b") // if -go = 1.16, ""
	_ = err
}
