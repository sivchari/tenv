package tenv

import (
	"go/ast"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"golang.org/x/mod/modfile"
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
	F     = "force"
	fflag bool
	A     = "all"
	aflag bool
)

func init() {
	Analyzer.Flags.BoolVar(&fflag, F, false, "the force option will also run against code prior to Go1.17")
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
					switch decl := decl.(type) {
					case *ast.FuncDecl:
						checkFunc(pass, decl)
					case *ast.GenDecl:
						if aflag {
							checkGenDecl(pass, decl)
						}
					}
				}
			}
		}
	})

	return nil, nil
}

func checkFunc(pass *analysis.Pass, n *ast.FuncDecl) {
	if targetRunner(n) {
		for _, stmt := range n.Body.List {
			switch stmt := stmt.(type) {
			case *ast.ExprStmt:
				if !checkExprStmt(pass, stmt, n) {
					continue
				}
			case *ast.IfStmt:
				if !checkIfStmt(pass, stmt, n) {
					continue
				}
			case *ast.AssignStmt:
				if !checkAssignStmt(pass, stmt, n) {
					continue
				}
			}
		}
	}
}

func checkExprStmt(pass *analysis.Pass, stmt *ast.ExprStmt, n *ast.FuncDecl) bool {
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
	funName := x.Name + "." + fun.Sel.Name
	if funName == "os.Setenv" {
		if checkVersion() {
			pass.Reportf(stmt.Pos(), "func %s is not using testing.Setenv", n.Name.Name)
		}
		if isForceExec() {
			pass.Reportf(stmt.Pos(), "func %s is not using testing.Setenv", n.Name.Name)
		}
	}
	return true
}

func checkIfStmt(pass *analysis.Pass, stmt *ast.IfStmt, n *ast.FuncDecl) bool {
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
	funName := x.Name + "." + fun.Sel.Name
	if funName == "os.Setenv" {
		if checkVersion() {
			pass.Reportf(stmt.Pos(), "func %s is not using testing.Setenv", n.Name.Name)
		}
		if isForceExec() {
			pass.Reportf(stmt.Pos(), "func %s is not using testing.Setenv", n.Name.Name)
		}
	}
	return true
}

func checkAssignStmt(pass *analysis.Pass, stmt *ast.AssignStmt, n *ast.FuncDecl) bool {
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
	funName := x.Name + "." + fun.Sel.Name
	if funName == "os.Setenv" {
		if checkVersion() {
			pass.Reportf(stmt.Pos(), "func %s is not using testing.Setenv", n.Name.Name)
		}
		if isForceExec() {
			pass.Reportf(stmt.Pos(), "func %s is not using testing.Setenv", n.Name.Name)
		}
	}
	return true
}

func checkGenDecl(pass *analysis.Pass, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		if len(valueSpec.Values) == 0 {
			continue
		}
		callExpr, ok := valueSpec.Values[0].(*ast.CallExpr)
		if !ok {
			continue
		}
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		x, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			continue
		}
		variable := valueSpec.Names[0].Name
		funName := x.Name + "." + selectorExpr.Sel.Name
		if funName == "os.Setenv" {
			if checkVersion() {
				pass.Reportf(valueSpec.Pos(), "variable %s is not using testing.Setenv", variable)
			}
			if isForceExec() {
				pass.Reportf(valueSpec.Pos(), "variable %s is not using testing.Setenv", variable)
			}
		}
	}
}

func checkVersion() bool {
	data, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Printf("read go.mod error: %v", err)
		return false
	}
	mod, err := modfile.Parse("", data, nil)
	if err != nil {
		log.Printf("parse go.mod error: %v", err)
		return false
	}
	floatVersion, err := strconv.ParseFloat(mod.Go.Version, 64)
	if err != nil {
		log.Printf("parse go.mod version error: %v", err)
		return false
	}
	return floatVersion >= 1.17
}

func isForceExec() bool {
	return fflag
}

func targetRunner(funcDecl *ast.FuncDecl) bool {
	if aflag {
		return true
	}
	params := funcDecl.Type.Params.List
	for _, p := range params {
		switch typ := p.Type.(type) {
		case *ast.StarExpr:
			if checkStarExprTarget(typ) {
				return true
			}
		case *ast.SelectorExpr:
			if checkSelectorExprTarget(typ) {
				return true
			}
		}
	}
	return false
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
		return false
	}
}

func checkSelectorExprTarget(typ *ast.SelectorExpr) bool {
	x, ok := typ.X.(*ast.Ident)
	if !ok {
		return false
	}
	targetName := x.Name + "." + typ.Sel.Name
	return targetName == "testing.TB"
}
