package gofileparser

// ParseGoFile parses a Go source file and returns a GFP_GoFile structure.
//
// Parameters:
//   - filePath: string - The path to the Go source file to be parsed.
//
// Returns:
//   - *GFP_GoFile: A pointer to the parsed file structure.
//   - error: Any error encountered during parsing.
//
// This function is the main entry point for parsing a Go file. It reads the file,
// parses its contents, and returns a structured representation of the Go file.
// If any error occurs during file reading or parsing, it returns nil and the error.
func ParseGoFile(filePath string) (*GFPGoFile, error) {
	return parseGoFile(filePath)
}

// ParseGoPackage parses all Go files in a directory and returns a slice of GFP_GoFile structures.
//
// Parameters:
//   - dirPath: string - The path to the directory containing Go files.
//
// Returns:
//   - []*GFP_GoFile: A slice of pointers to the parsed file structures.
//   - error: Any error encountered during parsing.
//
// This function parses all .go files in the specified directory, excluding test files.
// It returns a slice of parsed file structures and any error encountered during the process.
func ParseGoPackage(dirPath string) ([]*GFPGoFile, error) {
	return parseGoPackage(dirPath)
}
