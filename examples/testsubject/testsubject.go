// Package testsubject is a sample package for testing the Go file parser.
package testsubject

import (
	"fmt"
	"strings"
)

// TestConstant is a sample constant
const TestConstant = "This is a test constant"

// AnotherConstant is another sample constant with a type
const AnotherConstant string = "This is another test constant"

// TestVariable is a sample variable
var TestVariable int

// AnotherVariable is another sample variable with an initial value
var AnotherVariable = "This is a test variable"

// TestStruct is a sample struct
type TestStruct struct {
	Field1 string
	Field2 int
}

// TestInterface is a sample interface
type TestInterface interface {
	Method1() string
	Method2(int) error
}

// TestFunction is a sample function
func TestFunction(param1 string, param2 int) (string, error) {
	return fmt.Sprintf("Param1: %s, Param2: %d", param1, param2), nil
}

// TestMethod is a sample method for TestStruct
func (t *TestStruct) TestMethod(input string) string {
	return strings.ToUpper(input)
}

// This is a standalone comment that should be captured

/*
This is a multi-line comment
that should also be captured
*/
