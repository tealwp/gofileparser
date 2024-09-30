package gofileparser

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExprToString(t *testing.T) {
	tests := []struct {
		name     string
		expr     ast.Expr
		expected string
	}{
		{
			name:     "Nil expression",
			expr:     nil,
			expected: "",
		},
		{
			name:     "Basic literal",
			expr:     &ast.BasicLit{Kind: token.INT, Value: "42"},
			expected: "42",
		},
		{
			name: "Binary expression",
			expr: &ast.BinaryExpr{
				X:  &ast.BasicLit{Kind: token.INT, Value: "1"},
				Op: token.ADD,
				Y:  &ast.BasicLit{Kind: token.INT, Value: "2"},
			},
			expected: "1 + 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := exprToString(tt.expr)
			if result != tt.expected {
				t.Errorf("exprToString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBlockStmtToString(t *testing.T) {
	tests := []struct {
		name     string
		block    *ast.BlockStmt
		expected string
	}{
		{
			name:     "Nil block",
			block:    nil,
			expected: "",
		},
		{
			name: "Empty block",
			block: &ast.BlockStmt{
				List: []ast.Stmt{},
			},
			expected: "{\n}",
		},
		{
			name: "Block with statement",
			block: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ExprStmt{
						X: &ast.BasicLit{Kind: token.INT, Value: "42"},
					},
				},
			},
			expected: "{\n42\n}", // Removed leading tab
		},
		{
			name: "Block with multiple statements",
			block: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ExprStmt{
						X: &ast.BasicLit{Kind: token.INT, Value: "1"},
					},
					&ast.ExprStmt{
						X: &ast.BasicLit{Kind: token.INT, Value: "2"},
					},
				},
			},
			expected: "{\n1\n2\n}", // Removed leading tabs
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := blockStmtToString(tt.block)
			if result != tt.expected {
				t.Errorf("blockStmtToString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsTestFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected bool
	}{
		{
			name:     "Test file",
			filePath: "example_test.go",
			expected: true,
		},
		{
			name:     "Non-test Go file",
			filePath: "example.go",
			expected: false,
		},
		{
			name:     "Non-Go file",
			filePath: "example.txt",
			expected: false,
		},
		{
			name:     "File with _test in the middle",
			filePath: "example_test_file.go",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isTestFile(tt.filePath)
			if result != tt.expected {
				t.Errorf("isTestFile() = %v, want %v", result, tt.expected)
			}
		})
	}
}
