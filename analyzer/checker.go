package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			ce, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := ce.Fun.(*ast.SelectorExpr)
			if !ok || !isLoggerMethod(sel.Sel.Name) {
				return true
			}
			obj := pass.TypesInfo.ObjectOf(sel.Sel)
			if obj == nil || obj.Pkg() == nil {
				return true
			}
			path := obj.Pkg().Path()
			if !isLoggerExpr(path) {
				return true
			}

			return true
		})
	}
	return nil, nil
}

func isLoggerExpr(path string) bool {
	switch path {
	case "log/slog", "go.uber.org/zap":
		return true
	default:
		return false
	}
}

func isLoggerMethod(name string) bool {
	switch name {
	case "Info", "Error", "Warn", "Debug":
		return true
	default:
		return false
	}
}
