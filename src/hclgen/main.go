package main

import (
	"fmt"
	"hclgen/helpers"
)

func main() {
	parser := helpers.Parser{WorkingDirectory: "/tmp/"}

	err := parser.Parse("./testOutputDir")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n\n", parser)
}
