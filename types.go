package gofileparser

// GFPGoFile represents the structure of a parsed Go file.
type GFPGoFile struct {
	Package    string          // Name of the package
	Imports    []GFPImport    // List of imports
	Constants  []GFPConstant  // List of constants
	Variables  []GFPVariable  // List of variables
	Types      []GFPType      // List of type definitions
	Functions  []GFPFunction  // List of functions
	Methods    []GFPMethod    // List of methods
	Interfaces []GFPInterface // List of interfaces
	Comments   []GFPComment   // List of comments not associated with declarations
	FileDoc    string          // File-level documentation comment
	Content    string          // Entire file content
}

// GFPImport represents a single import statement.
type GFPImport struct {
	Path string // Import path (e.g., "fmt")
	Name string // Local name (alias) for the import, if any
	Line int    // Line number where the import is declared
}

// GFPConstant represents a constant declaration.
type GFPConstant struct {
	Name  string // Name of the constant
	Type  string // Type of the constant (may be empty if inferred)
	Value string // Value of the constant
	Doc   string // Associated documentation comment
	Line  int    // Line number where the constant is declared
}

// GFPVariable represents a variable declaration.
type GFPVariable struct {
	Name  string // Name of the variable
	Type  string // Type of the variable (may be empty if inferred)
	Value string // Initial value of the variable (may be empty)
	Doc   string // Associated documentation comment
	Line  int    // Line number where the variable is declared
}

// GFPType represents a type definition.
type GFPType struct {
	Name string // Name of the type
	Def  string // Definition of the type
	Doc  string // Associated documentation comment
	Line int    // Line number where the type is declared
}

// GFPFunction represents a function declaration.
type GFPFunction struct {
	Name       string          // Name of the function
	Parameters []GFPParameter // List of parameters
	ReturnType string          // Return type(s)
	Body       string          // Function body
	Doc        string          // Associated documentation comment
	Line       int             // Line number where the function is declared
}

// GFPMethod represents a method declaration.
type GFPMethod struct {
	Receiver   string          // Receiver type
	Name       string          // Name of the method
	Parameters []GFPParameter // List of parameters
	ReturnType string          // Return type(s)
	Body       string          // Method body
	Doc        string          // Associated documentation comment
	Line       int             // Line number where the method is declared
}

// GFPInterface represents an interface declaration.
type GFPInterface struct {
	Name    string                // Name of the interface
	Methods []GFPInterfaceMethod // List of methods in the interface
	Doc     string                // Associated documentation comment
	Line    int                   // Line number where the interface is declared
}

// GFPInterfaceMethod represents a method in an interface declaration.
type GFPInterfaceMethod struct {
	Name       string          // Name of the method
	Parameters []GFPParameter // List of parameters
	ReturnType string          // Return type(s)
	Line       int             // Line number where the interface method is declared
}

// GFPParameter represents a function or method parameter.
type GFPParameter struct {
	Name string // Name of the parameter (may be empty for unnamed parameters)
	Type string // Type of the parameter
}

// GFPComment represents a comment in the Go file.
type GFPComment struct {
	Text string // Text of the comment
	Line int    // Line number where the comment appears
}
