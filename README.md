### GoFileParser

gofileparser is a simple wrapper around the following go standard tooling packages:
- "go/ast"
- "go/parser"
- "go/token"

The intended use is internal, and acts as the parser for other projects of mine.

### Usage

```go
package main

import (
	"fmt"

	"github.com/tealwp/gofileparser"
)

func main() {
	// Parse a file
	file, err := gofileparser.ParseFile("path/to/file.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Print the file contents
	fmt.Println(file.Contents)
}
```

### Features

* Parses files into a `File` struct
* Provides access to the file contents

### Installation

```bash
go get github.com/tealwp/gofileparser
```

### Examples

```bash
go run examples/example.go
```

### Needed Fixes

- Type definitions aren't getting their documentation (comments) added.... [The problem seems to be a bit deeper than expected.](https://stackoverflow.com/questions/19580688/go-parser-not-detecting-doc-comments-on-struct-type)


