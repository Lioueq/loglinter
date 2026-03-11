package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			be, ok := n.(*ast.BinaryExpr)
			if !ok {
				return true
			}
			if be.Op == token.ADD {
				pass.Reportf(be.Pos(), "addiction found")
				return true
			}
			return true
		})
	}
	return nil, nil
}
