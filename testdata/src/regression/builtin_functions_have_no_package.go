package regression_test

import "testing"

func TestBuiltInFunctionsHaveNoPkg(t *testing.T) {
	// Built-in definitions have no package, which means that "go/types".Object.Pkg() returns nil,
	// which in turn panics when its method Name() is called.
	var items []int
	items = append(items, 0)
}
