package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

var knownMethods = map[string]struct{}{
	"Info": {}, "Error": {}, "Warn": {}, "Debug": {}, "Fatal": {},
	"Infow": {}, "Errorw": {}, "Warnw": {}, "Debugw": {}, "Fatalw": {},
	"Infof": {}, "Errorf": {}, "Warnf": {}, "Debugf": {}, "Fatalf": {},
}

func run(pass *analysis.Pass) (any, error) {
	sensitivePattern, err := buildSensitivePattern()
	if err != nil {
		return nil, fmt.Errorf("loglinter: invalid sensitive pattern in config: %w", err)
	}
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			ce, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := ce.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			if _, ok := knownMethods[sel.Sel.Name]; !ok {
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
			if len(ce.Args) == 0 {
				return true
			}
			checkArg(pass, ce.Args[0], sensitivePattern)
			return true
		})
	}
	return nil, nil
}

func checkArg(pass *analysis.Pass, firstArg ast.Expr, sensitivePattern *regexp.Regexp) {
	switch arg := firstArg.(type) {
	case *ast.BasicLit:
		checkLiteralArg(pass, arg)
	case *ast.BinaryExpr:
		checkConcatArg(pass, arg, sensitivePattern)
	}
}

func checkLiteralArg(pass *analysis.Pass, arg *ast.BasicLit) {
	if arg.Kind != token.STRING {
		return
	}
	msg, err := strconv.Unquote(arg.Value)
	if err != nil {
		return
	}
	checkLowercase(pass, arg.Pos(), msg)
	checkEnglishLetters(pass, arg.Pos(), msg)
	checkSpecialChars(pass, arg.Pos(), msg)
}

func checkConcatArg(pass *analysis.Pass, expr *ast.BinaryExpr, sensitivePattern *regexp.Regexp) {
	for i, lit := range extractLiterals(expr) {
		if i == 0 {
			checkLowercase(pass, expr.Pos(), lit)
		}
		checkEnglishLetters(pass, expr.Pos(), lit)
		checkSpecialChars(pass, expr.Pos(), lit)
	}
	checkSensitiveData(pass, expr, sensitivePattern)
}

func checkLowercase(pass *analysis.Pass, pos token.Pos, msg string) {
	r, _ := utf8.DecodeRuneInString(msg)
	if unicode.IsUpper(r) {
		pass.Reportf(pos, "log message must start with a lowercase letter")
	}
}

func checkEnglishLetters(pass *analysis.Pass, pos token.Pos, msg string) {
	for _, r := range msg {
		if r > unicode.MaxASCII && unicode.IsLetter(r) {
			pass.Reportf(pos, "log message must be in English")
			return
		}
	}
}

func checkSpecialChars(pass *analysis.Pass, pos token.Pos, msg string) {
	for _, r := range msg {
		if r > unicode.MaxASCII && unicode.IsLetter(r) {
			continue
		}
		if !isAllowedRune(r) {
			pass.Reportf(pos, "log message must not contain special characters or emoji")
			return
		}
	}
}

func isAllowedRune(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9') ||
		r == ' ' || r == '-' || r == '_' || r == '=' || r == ',' || r == '%' || r == ':'
}

func checkSensitiveData(pass *analysis.Pass, expr *ast.BinaryExpr, sensitivePattern *regexp.Regexp) {
	if sensitivePattern == nil || !containsIdent(expr) {
		return
	}
	if slices.ContainsFunc(extractIdents(expr), sensitivePattern.MatchString) {
		pass.Reportf(expr.Pos(), "log message may contain sensitive data")
		return
	}
}

func extractLiterals(expr ast.Expr) []string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			val, err := strconv.Unquote(e.Value)
			if err == nil {
				return []string{val}
			}
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			return append(extractLiterals(e.X), extractLiterals(e.Y)...)
		}
	}
	return nil
}

func extractIdents(expr ast.Expr) []string {
	switch e := expr.(type) {
	case *ast.Ident:
		return []string{strings.ToLower(e.Name)}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			return append(extractIdents(e.X), extractIdents(e.Y)...)
		}
	}
	return nil
}

func containsIdent(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.Ident:
		return true
	case *ast.BinaryExpr:
		return containsIdent(e.X) || containsIdent(e.Y)
	}
	return false
}

func isLoggerExpr(path string) bool {
	switch path {
	case "log/slog", "go.uber.org/zap":
		return true
	default:
		return false
	}
}
