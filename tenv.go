package tenv

import (
	"go/ast"
	"log"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "tenv is analyzer that detects environment variable not using t.Setenv"

// Analyzer is tenv analyzer
var Analyzer = &analysis.Analyzer{
	Name: "tenv",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var F bool

func init() {
	Analyzer.Flags.BoolVar(&F, "f", false, "the force option will also run against code prior to Go1.17")
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
						checkGenDecl(pass, decl)
					}
				}
			}
		}
	})

	return nil, nil
}

func checkFunc(pass *analysis.Pass, n *ast.FuncDecl) {
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
		foldV := checkVersion()
		if foldV >= 1.17 || isForceExec() {
			pass.Reportf(stmt.Pos(), "func %s is not using t.Setenv", n.Name.Name)
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
		foldV := checkVersion()
		if foldV >= 1.17 || isForceExec() {
			pass.Reportf(stmt.Pos(), "func %s is not using t.Setenv", n.Name.Name)
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
		foldV := checkVersion()
		if foldV >= 1.17 || isForceExec() {
			pass.Reportf(stmt.Pos(), "func %s is not using t.Setenv", n.Name.Name)
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
			foldV := checkVersion()
			if foldV >= 1.17 || isForceExec() {
				pass.Reportf(valueSpec.Pos(), "variable %s is not using t.Setenv", variable)
			}
		}
	}
}

func checkVersion() float64 {
	version := strings.Trim(runtime.Version(), "go")
	foldV, err := strconv.ParseFloat(version[0:4], 64)
	if err != nil {
		log.Println(err)
	}
	return foldV
}

func isForceExec() bool {
	return F
}
