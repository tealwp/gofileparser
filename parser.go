package gofileparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// parseGoFile parses a Go source file and returns a GFP_GoFile structure.
// This is the internal implementation of ParseGoFile.
func parseGoFile(filePath string) (*GFP_GoFile, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, content, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	goFile := &GFP_GoFile{}

	goFile.Package = file.Name.Name
	goFile.Content = string(content)

	if file.Doc != nil {
		goFile.FileDoc = file.Doc.Text()
	}

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.IMPORT:
				goFile.Imports = append(goFile.Imports, parseImports(fset, d)...)
			case token.CONST:
				goFile.Constants = append(goFile.Constants, parseConstants(fset, d)...)
			case token.VAR:
				goFile.Variables = append(goFile.Variables, parseVariables(fset, d)...)
			case token.TYPE:
				parseTypes(fset, d, goFile)
			}
		case *ast.FuncDecl:
			if d.Recv == nil {
				goFile.Functions = append(goFile.Functions, parseFunction(fset, d))
			} else {
				goFile.Methods = append(goFile.Methods, parseMethod(fset, d))
			}
		}
	}

	goFile.Comments = parseComments(fset, file)

	return goFile, nil
}

// parseGoPackage parses all Go files in a directory and returns a slice of GFP_GoFile structures.
func parseGoPackage(dirPath string) ([]*GFP_GoFile, error) {
	files, err := filepath.Glob(filepath.Join(dirPath, "*.go"))
	if err != nil {
		return nil, fmt.Errorf("error finding Go files: %w", err)
	}

	var parsedFiles []*GFP_GoFile
	for _, file := range files {
		// Skip test files
		if filepath.Ext(file) == ".go" && !isTestFile(file) {
			parsedFile, err := ParseGoFile(file)
			if err != nil {
				return nil, fmt.Errorf("error parsing file %s: %w", file, err)
			}
			parsedFiles = append(parsedFiles, parsedFile)
		}
	}

	return parsedFiles, nil
}

// parseImports extracts import declarations from a GenDecl.
func parseImports(fset *token.FileSet, decl *ast.GenDecl) []GFP_Import {
	var imports []GFP_Import
	for _, spec := range decl.Specs {
		if is, ok := spec.(*ast.ImportSpec); ok {
			imp := GFP_Import{
				Path: is.Path.Value,
				Line: fset.Position(is.Pos()).Line,
			}
			if is.Name != nil {
				imp.Name = is.Name.Name
			}
			imports = append(imports, imp)
		}
	}
	return imports
}

// parseConstants extracts constant declarations from a GenDecl.
func parseConstants(fset *token.FileSet, decl *ast.GenDecl) []GFP_Constant {
	var constants []GFP_Constant
	for _, spec := range decl.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			for i, name := range vs.Names {
				c := GFP_Constant{
					Name: name.Name,
					Type: exprToString(vs.Type),
					Doc:  vs.Doc.Text(),
					Line: fset.Position(name.Pos()).Line,
				}
				if i < len(vs.Values) {
					c.Value = exprToString(vs.Values[i])
				}
				constants = append(constants, c)
			}
		}
	}
	return constants
}

// parseVariables extracts variable declarations from a GenDecl.
func parseVariables(fset *token.FileSet, decl *ast.GenDecl) []GFP_Variable {
	var variables []GFP_Variable
	for _, spec := range decl.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			for i, name := range vs.Names {
				v := GFP_Variable{
					Name: name.Name,
					Type: exprToString(vs.Type),
					Doc:  vs.Doc.Text(),
					Line: fset.Position(name.Pos()).Line,
				}
				if i < len(vs.Values) {
					v.Value = exprToString(vs.Values[i])
				}
				variables = append(variables, v)
			}
		}
	}
	return variables
}

// parseTypes extracts type declarations from a GenDecl and adds them to the GFP_GoFile.
func parseTypes(fset *token.FileSet, decl *ast.GenDecl, goFile *GFP_GoFile) {
	for _, spec := range decl.Specs {
		if ts, ok := spec.(*ast.TypeSpec); ok {
			if _, ok := ts.Type.(*ast.InterfaceType); ok {
				goFile.Interfaces = append(goFile.Interfaces, parseInterface(fset, ts))
			} else {
				goFile.Types = append(goFile.Types, parseType(fset, ts, decl))
			}
		}
	}
}

// parseType extracts a single type definition from a TypeSpec.
func parseType(fset *token.FileSet, ts *ast.TypeSpec, decl *ast.GenDecl) GFP_Type {
	return GFP_Type{
		Name: ts.Name.Name,
		Def:  exprToString(ts.Type),
		Doc:  decl.Doc.Text(),
		Line: fset.Position(ts.Name.Pos()).Line,
	}
}

// parseInterface extracts an interface definition from a TypeSpec.
func parseInterface(fset *token.FileSet, ts *ast.TypeSpec) GFP_Interface {
	iface := GFP_Interface{
		Name: ts.Name.Name,
		Doc:  ts.Doc.Text(),
		Line: fset.Position(ts.Name.Pos()).Line,
	}
	if it, ok := ts.Type.(*ast.InterfaceType); ok {
		for _, method := range it.Methods.List {
			iface.Methods = append(iface.Methods, parseInterfaceMethod(fset, method))
		}
	}
	return iface
}

// parseInterfaceMethod extracts a method definition from an interface field.
func parseInterfaceMethod(fset *token.FileSet, field *ast.Field) GFP_InterfaceMethod {
	method := GFP_InterfaceMethod{
		Name:       field.Names[0].Name,
		Parameters: parseParameters(field.Type.(*ast.FuncType).Params),
		ReturnType: parseReturnType(field.Type.(*ast.FuncType).Results),
		Line:       fset.Position(field.Names[0].Pos()).Line,
	}
	return method
}

// parseFunction extracts a function definition from a FuncDecl.
func parseFunction(fset *token.FileSet, decl *ast.FuncDecl) GFP_Function {
	return GFP_Function{
		Name:       decl.Name.Name,
		Parameters: parseParameters(decl.Type.Params),
		ReturnType: parseReturnType(decl.Type.Results),
		Body:       blockStmtToString(decl.Body),
		Doc:        decl.Doc.Text(),
		Line:       fset.Position(decl.Name.Pos()).Line,
	}
}

// parseMethod extracts a method definition from a FuncDecl.
func parseMethod(fset *token.FileSet, decl *ast.FuncDecl) GFP_Method {
	return GFP_Method{
		Receiver:   exprToString(decl.Recv.List[0].Type),
		Name:       decl.Name.Name,
		Parameters: parseParameters(decl.Type.Params),
		ReturnType: parseReturnType(decl.Type.Results),
		Body:       blockStmtToString(decl.Body),
		Doc:        decl.Doc.Text(),
		Line:       fset.Position(decl.Name.Pos()).Line,
	}
}

// parseParameters extracts parameter definitions from a FieldList.
func parseParameters(fields *ast.FieldList) []GFP_Parameter {
	var params []GFP_Parameter
	if fields != nil {
		for _, field := range fields.List {
			for _, name := range field.Names {
				params = append(params, GFP_Parameter{
					Name: name.Name,
					Type: exprToString(field.Type),
				})
			}
		}
	}
	return params
}

// parseReturnType extracts return type(s) from a FieldList.
func parseReturnType(fields *ast.FieldList) string {
	if fields == nil || len(fields.List) == 0 {
		return ""
	}
	var types []string
	for _, field := range fields.List {
		types = append(types, exprToString(field.Type))
	}
	if len(types) == 1 {
		return types[0]
	}
	return "(" + strings.Join(types, ", ") + ")"
}

// parseComments extracts comments from a File that are not associated with declarations.
func parseComments(fset *token.FileSet, file *ast.File) []GFP_Comment {
	var comments []GFP_Comment
	for _, commentGroup := range file.Comments {
		for _, comment := range commentGroup.List {
			comments = append(comments, GFP_Comment{
				Text: comment.Text,
				Line: fset.Position(comment.Pos()).Line,
			})
		}
	}
	return comments
}
