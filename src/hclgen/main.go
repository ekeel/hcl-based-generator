package main

import (
	"flag"
	"fmt"
	"hclgen/helpers"
)

func main() {
	var workingDirectory string
	var outputDirectory string

	flag.StringVar(&workingDirectory, "i", "./", "The directory to read the .hcl files from.")
	flag.StringVar(&outputDirectory, "o", "", "The directory to output the rendered files to.")

	flag.Parse()

	parser := helpers.Parser{WorkingDirectory: workingDirectory}

	err := parser.Parse(outputDirectory)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n\n", parser)
}
