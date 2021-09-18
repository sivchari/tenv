package tenv

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "tenv is analyzer that detects using os.Setenv instead of t.Setenv since Go1.17"

// Analyzer is tenv analyzer
var Analyzer = &analysis.Analyzer{
	Name: "tenv",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var (
	A     = "all"
	aflag bool
)

func init() {
	Analyzer.Flags.BoolVar(&aflag, A, false, "the all option will run against all method in test file")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			if strings.HasSuffix(pass.Fset.File(n.Pos()).Name(), "_test.go") {
				for _, decl := range n.Decls {
					funcDecl, ok := decl.(*ast.FuncDecl)
					if !ok {
						continue
					}
					checkFunc(pass, funcDecl)
				}
			}
		}
	})

	return nil, nil
}

func checkFunc(pass *analysis.Pass, n *ast.FuncDecl) {
	argName, ok := targetRunner(n)
	if ok {
		for _, stmt := range n.Body.List {
			switch stmt := stmt.(type) {
			case *ast.ExprStmt:
				if !checkExprStmt(pass, stmt, n, argName) {
					continue
				}
			case *ast.IfStmt:
				if !checkIfStmt(pass, stmt, n, argName) {
					continue
				}
			case *ast.AssignStmt:
				if !checkAssignStmt(pass, stmt, n, argName) {
					continue
				}
			}
		}
	}
}

func checkExprStmt(pass *analysis.Pass, stmt *ast.ExprStmt, n *ast.FuncDecl, argName string) bool {
	callExpr, ok := stmt.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	fun, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	x, ok := fun.X.(*ast.Ident)
	if !ok {
		return false
	}
	targetName := x.Name + "." + fun.Sel.Name
	if targetName == "os.Setenv" {
		if argName == "" {
			argName = "testing"
		}
		pass.Reportf(stmt.Pos(), "os.Setenv() can be replaced by `%s.Setenv()` in %s", argName, n.Name.Name)
	}
	return true
}

func checkIfStmt(pass *analysis.Pass, stmt *ast.IfStmt, n *ast.FuncDecl, argName string) bool {
	assignStmt, ok := stmt.Init.(*ast.AssignStmt)
	if !ok {
		return false
	}
	rhs, ok := assignStmt.Rhs[0].(*ast.CallExpr)
	if !ok {
		return false
	}
	fun, ok := rhs.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	x, ok := fun.X.(*ast.Ident)
	if !ok {
		return false
	}
	targetName := x.Name + "." + fun.Sel.Name
	if targetName == "os.Setenv" {
		if argName == "" {
			argName = "testing"
		}
		pass.Reportf(stmt.Pos(), "os.Setenv() can be replaced by `%s.Setenv()` in %s", argName, n.Name.Name)
	}
	return true
}

func checkAssignStmt(pass *analysis.Pass, stmt *ast.AssignStmt, n *ast.FuncDecl, argName string) bool {
	rhs, ok := stmt.Rhs[0].(*ast.CallExpr)
	if !ok {
		return false
	}
	fun, ok := rhs.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	x, ok := fun.X.(*ast.Ident)
	if !ok {
		return false
	}
	targetName := x.Name + "." + fun.Sel.Name
	if targetName == "os.Setenv" {
		if argName == "" {
			argName = "testing"
		}
		pass.Reportf(stmt.Pos(), "os.Setenv() can be replaced by `%s.Setenv()` in %s", argName, n.Name.Name)
	}
	return true
}

func targetRunner(funcDecl *ast.FuncDecl) (string, bool) {
	params := funcDecl.Type.Params.List
	for _, p := range params {
		switch typ := p.Type.(type) {
		case *ast.StarExpr:
			if checkStarExprTarget(typ) {
				argName := p.Names[0].Name
				return argName, true
			}
		case *ast.SelectorExpr:
			if checkSelectorExprTarget(typ) {
				argName := p.Names[0].Name
				return argName, true
			}
		}
	}
	return "", false
}

func checkStarExprTarget(typ *ast.StarExpr) bool {
	selector, ok := typ.X.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	x, ok := selector.X.(*ast.Ident)
	if !ok {
		return false
	}
	targetName := x.Name + "." + selector.Sel.Name
	switch targetName {
	case "testing.T", "testing.B":
		return true
	default:
		if aflag {
			return true
		}
		return false
	}
}

func checkSelectorExprTarget(typ *ast.SelectorExpr) bool {
	x, ok := typ.X.(*ast.Ident)
	if !ok {
		return false
	}
	targetName := x.Name + "." + typ.Sel.Name
	if targetName == "testing.TB" {
		return true
	}
	if aflag {
		return true
	}
	return false
}
