package gofileparser

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"path/filepath"
	"strings"
)

// exprToString converts an ast.Expr to its string representation.
//
// Parameters:
//   - expr: ast.Expr - The expression to convert to a string.
//
// Returns:
//   - string: The string representation of the expression.
//
// This function takes an AST expression and converts it to a formatted Go code string.
// If the expression is nil or there's an error in formatting, it returns an empty string.
func exprToString(expr ast.Expr) string {
	if expr == nil {
		return ""
	}
	var buf bytes.Buffer
	if err := format.Node(&buf, token.NewFileSet(), expr); err != nil {
		return ""
	}
	return buf.String()
}

// blockStmtToString converts an ast.BlockStmt to its string representation.
//
// Parameters:
//   - block: *ast.BlockStmt - The block statement to convert to a string.
//
// Returns:
//   - string: The string representation of the block statement.
//
// This function takes an AST block statement (typically a function body) and
// converts it to a formatted Go code string. If the block is nil, it returns an empty string.
func blockStmtToString(block *ast.BlockStmt) string {
	if block == nil {
		return ""
	}
	var buf bytes.Buffer
	format.Node(&buf, token.NewFileSet(), block)
	content := buf.String()
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimPrefix(line, "\t")
	}
	return strings.Join(lines, "\n")
}

// isTestFile checks if a file is a Go test file.
//
// Parameters:
//   - filePath: string - The path to the file to check.
//
// Returns:
//   - bool: true if the file is a test file, false otherwise.
//
// This function checks if a file name ends with "_test.go", which is the
// convention for Go test files.
func isTestFile(filePath string) bool {
	return filepath.Ext(filePath) == ".go" && filepath.Base(filePath)[len(filepath.Base(filePath))-8:] == "_test.go"
}
