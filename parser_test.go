package gofileparser

import (
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func TestParseGoFile(t *testing.T) {
	// Create a temporary Go file for testing
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.go")
	content := []byte(`
package main

import "fmt"

const greeting = "Hello, World!"

var message string

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(greeting)
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}
`)
	err := os.WriteFile(tempFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Parse the file
	goFile, err := parseGoFile(tempFile)
	if err != nil {
		t.Fatalf("ParseGoFile failed: %v", err)
	}

	// Check parsed content
	if goFile.Package != "main" {
		t.Errorf("Expected package 'main', got '%s'", goFile.Package)
	}

	if len(goFile.Imports) != 1 || goFile.Imports[0].Path != "\"fmt\"" {
		t.Errorf("Import not parsed correctly")
	}

	if len(goFile.Constants) != 1 || goFile.Constants[0].Name != "greeting" {
		t.Errorf("Constant not parsed correctly")
	}

	if len(goFile.Variables) != 1 || goFile.Variables[0].Name != "message" {
		t.Errorf("Variable not parsed correctly")
	}

	if len(goFile.Types) != 1 || goFile.Types[0].Name != "Person" {
		t.Errorf("Type not parsed correctly")
	}

	if len(goFile.Functions) != 1 || goFile.Functions[0].Name != "main" {
		t.Errorf("Function not parsed correctly")
	}

	if len(goFile.Methods) != 1 || goFile.Methods[0].Name != "SayHello" {
		t.Errorf("Method not parsed correctly")
	}

	if goFile.Content != string(content) {
		t.Errorf("File content not parsed correctly")
	}
}

func TestParseGoPackage(t *testing.T) {
	// Create a temporary directory with multiple Go files
	tempDir := t.TempDir()
	createTempGoFile(t, tempDir, "file1.go", "package main\n\nfunc Func1() {}\n")
	createTempGoFile(t, tempDir, "file2.go", "package main\n\nfunc Func2() {}\n")
	createTempGoFile(t, tempDir, "file_test.go", "package main\n\nfunc TestFunc() {}\n")

	// Parse the package
	files, err := parseGoPackage(tempDir)
	if err != nil {
		t.Fatalf("ParseGoPackage failed: %v", err)
	}

	// Check parsed content
	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}

	funcNames := make(map[string]bool)
	for _, file := range files {
		for _, fn := range file.Functions {
			funcNames[fn.Name] = true
		}
	}

	if !funcNames["Func1"] || !funcNames["Func2"] {
		t.Errorf("Expected functions Func1 and Func2, got %v", funcNames)
	}

	if funcNames["TestFunc"] {
		t.Errorf("Test file should not be parsed")
	}
}

func TestParseImports(t *testing.T) {
	fset := token.NewFileSet()
	importDecl := &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{Path: &ast.BasicLit{Value: "\"fmt\""}},
			&ast.ImportSpec{Name: ast.NewIdent("alias"), Path: &ast.BasicLit{Value: "\"some/package\""}},
		},
	}

	imports := parseImports(fset, importDecl)

	if len(imports) != 2 {
		t.Errorf("Expected 2 imports, got %d", len(imports))
	}

	if imports[0].Path != "\"fmt\"" || imports[0].Name != "" {
		t.Errorf("First import not parsed correctly")
	}

	if imports[1].Path != "\"some/package\"" || imports[1].Name != "alias" {
		t.Errorf("Second import not parsed correctly")
	}
}

func TestParseConstants(t *testing.T) {
	fset := token.NewFileSet()
	constDecl := &ast.GenDecl{
		Tok: token.CONST,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names:  []*ast.Ident{ast.NewIdent("Pi")},
				Type:   ast.NewIdent("float64"),
				Values: []ast.Expr{&ast.BasicLit{Kind: token.FLOAT, Value: "3.14"}},
			},
		},
	}

	constants := parseConstants(fset, constDecl)

	if len(constants) != 1 {
		t.Errorf("Expected 1 constant, got %d", len(constants))
	}

	if constants[0].Name != "Pi" || constants[0].Type != "float64" || constants[0].Value != "3.14" {
		t.Errorf("Constant not parsed correctly")
	}
}

func TestParseVariables(t *testing.T) {
	fset := token.NewFileSet()
	varDecl := &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{ast.NewIdent("message")},
				Type:  ast.NewIdent("string"),
			},
		},
	}

	variables := parseVariables(fset, varDecl)

	if len(variables) != 1 {
		t.Errorf("Expected 1 variable, got %d", len(variables))
	}

	if variables[0].Name != "message" || variables[0].Type != "string" {
		t.Errorf("Variable not parsed correctly")
	}
}

func TestParseFunction(t *testing.T) {
	fset := token.NewFileSet()
	funcDecl := &ast.FuncDecl{
		Name: ast.NewIdent("TestFunc"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{ast.NewIdent("x")}, Type: ast.NewIdent("int")},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{Type: ast.NewIdent("string")},
				},
			},
		},
		Body: &ast.BlockStmt{},
	}

	function := parseFunction(fset, funcDecl)

	if function.Name != "TestFunc" {
		t.Errorf("Expected function name TestFunc, got %s", function.Name)
	}

	if len(function.Parameters) != 1 || function.Parameters[0].Name != "x" || function.Parameters[0].Type != "int" {
		t.Errorf("Function parameters not parsed correctly")
	}

	if function.ReturnType != "string" {
		t.Errorf("Function return type not parsed correctly, got %s", function.ReturnType)
	}
}

func createTempGoFile(t *testing.T, dir, name, content string) {
	t.Helper()
	err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file %s: %v", name, err)
	}
}
