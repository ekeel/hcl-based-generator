package helpers

import (
	"bytes"
	"fmt"
	"hclgen/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

// Parser is a struct for reading, parsing, and rendering a project.
type Parser struct {
	WorkingDirectory string
	DiscoveredFiles  []string
	RawContent       string
	Project          model.Project
}

// Parse discoveries all the .hcl files in the working directory and parses them.
func (parser *Parser) Parse(outputDirectory string) (err error) {
	if _, err := os.Stat(parser.WorkingDirectory); os.IsNotExist(err) {
		return fmt.Errorf("error: the workding directory at %v does not exist", parser.WorkingDirectory)
	}

	err = parser.DiscoverFiles()
	if err != nil {
		return err
	}

	err = parser.ReadFiles()
	if err != nil {
		return err
	}

	err = parser.ParseContent()
	if err != nil {
		return err
	}

	err = parser.RenderTemplates(outputDirectory)
	if err != nil {
		return err
	}

	return nil
}

// DiscoverFiles walks the working directory and discovers all the .hcl files.
func (parser *Parser) DiscoverFiles() (err error) {
	parser.DiscoveredFiles, err = walkMatch(parser.WorkingDirectory, `(?m).*?\.hcl$`)
	if err != nil {
		return err
	}

	return nil
}

// ReadFiles reads all the discovered .hcl files to a single string for parsing.
func (parser *Parser) ReadFiles() (err error) {
	for _, file := range parser.DiscoveredFiles {
		rawContent, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		parser.RawContent = parser.RawContent + "\n\n\n" + string(rawContent)
	}

	return nil
}

// ParseContent parses the raw content read from the files to models.
func (parser *Parser) ParseContent() (err error) {
	parser.Project = model.Project{}

	hclParser := hclparse.NewParser()

	hclContent, parseDiags := hclParser.ParseHCL([]byte(parser.RawContent), "generatedFile")
	if parseDiags.HasErrors() {
		return err
	}

	decodeDiags := gohcl.DecodeBody(hclContent.Body, nil, &parser.Project)
	if decodeDiags.HasErrors() {
		return err
	}

	return nil
}

// RenderTemplates renders all of the files based on the project model.
func (parser *Parser) RenderTemplates(outputDirectory string) (err error) {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(outputDirectory, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, seed := range parser.Project.Seeds {
		outputPath := filepath.Join(outputDirectory, fmt.Sprintf("%s.groovy", seed.Name))

		err := seed.GenerateConstants()
		if err != nil {
			return err
		}

		var buff bytes.Buffer
		tpl := template.Must(template.ParseFiles("/home/codespace/workspace/hcl-based-generator/src/hclgen/template/seed.tpl"))
		tpl.Execute(&buff, *seed)

		file, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(buff.String())
		if err != nil {
			return err
		}
	}

	return nil
}
