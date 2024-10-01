package main

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/tealwp/gofileparser"
)

func main() {
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("----testing on example go file at ./examples/testsubject/testsubject.go----")
	fmt.Println("---------------------------------------------------------------------------")
	goFile, err := gofileparser.ParseGoFile("./examples/testsubject/testsubject.go")
	assert.NoError(nil, err)
	assert.NotNil(nil, goFile)

	fmt.Println(goFile)
}