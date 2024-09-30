package gofileparser

// GFP_GoFile represents the structure of a parsed Go file.
type GFP_GoFile struct {
	Package    string          // Name of the package
	Imports    []GFP_Import    // List of imports
	Constants  []GFP_Constant  // List of constants
	Variables  []GFP_Variable  // List of variables
	Types      []GFP_Type      // List of type definitions
	Functions  []GFP_Function  // List of functions
	Methods    []GFP_Method    // List of methods
	Interfaces []GFP_Interface // List of interfaces
	Comments   []GFP_Comment   // List of comments not associated with declarations
	FileDoc    string          // File-level documentation comment
}

// GFP_Import represents a single import statement.
type GFP_Import struct {
	Path string // Import path (e.g., "fmt")
	Name string // Local name (alias) for the import, if any
	Line int    // Line number where the import is declared
}

// GFP_Constant represents a constant declaration.
type GFP_Constant struct {
	Name  string // Name of the constant
	Type  string // Type of the constant (may be empty if inferred)
	Value string // Value of the constant
	Doc   string // Associated documentation comment
	Line  int    // Line number where the constant is declared
}

// GFP_Variable represents a variable declaration.
type GFP_Variable struct {
	Name  string // Name of the variable
	Type  string // Type of the variable (may be empty if inferred)
	Value string // Initial value of the variable (may be empty)
	Doc   string // Associated documentation comment
	Line  int    // Line number where the variable is declared
}

// GFP_Type represents a type definition.
type GFP_Type struct {
	Name string // Name of the type
	Def  string // Definition of the type
	Doc  string // Associated documentation comment
	Line int    // Line number where the type is declared
}

// GFP_Function represents a function declaration.
type GFP_Function struct {
	Name       string          // Name of the function
	Parameters []GFP_Parameter // List of parameters
	ReturnType string          // Return type(s)
	Body       string          // Function body
	Doc        string          // Associated documentation comment
	Line       int             // Line number where the function is declared
}

// GFP_Method represents a method declaration.
type GFP_Method struct {
	Receiver   string          // Receiver type
	Name       string          // Name of the method
	Parameters []GFP_Parameter // List of parameters
	ReturnType string          // Return type(s)
	Body       string          // Method body
	Doc        string          // Associated documentation comment
	Line       int             // Line number where the method is declared
}

// GFP_Interface represents an interface declaration.
type GFP_Interface struct {
	Name    string                // Name of the interface
	Methods []GFP_InterfaceMethod // List of methods in the interface
	Doc     string                // Associated documentation comment
	Line    int                   // Line number where the interface is declared
}

// GFP_InterfaceMethod represents a method in an interface declaration.
type GFP_InterfaceMethod struct {
	Name       string          // Name of the method
	Parameters []GFP_Parameter // List of parameters
	ReturnType string          // Return type(s)
	Line       int             // Line number where the interface method is declared
}

// GFP_Parameter represents a function or method parameter.
type GFP_Parameter struct {
	Name string // Name of the parameter (may be empty for unnamed parameters)
	Type string // Type of the parameter
}

// GFP_Comment represents a comment in the Go file.
type GFP_Comment struct {
	Text string // Text of the comment
	Line int    // Line number where the comment appears
}
