package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"teamcity2gitlab/pkg/convert"
)

func main() {

	inputDir := "../../xml"
	outputDir := "../../output"
	templateFile := "../../templates/gitlab-ci-java.go.tmpl"
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("Error reading input directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".xml") {
			continue
		}

		inputFile := filepath.Join(inputDir, file.Name())
		outputFile := filepath.Join(outputDir, strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))+".yml")

		convert.Convert(inputFile, outputFile, templateFile)

		log.Printf("Conversion completed. GitLab pipeline YAML data written to %s", outputFile)
	}
}
