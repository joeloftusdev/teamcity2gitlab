package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
	"teamcity2gitlab/pkg/convert"
)

type Config struct {
	Templates struct {
		JavaTemplate string `yaml:"java_template"`
	} `yaml:"templates"`
	TeamCity struct {
		DataDir string `yaml:"data_dir"`
	} `yaml:"teamcity"`
	GitLab struct {
		PipelineOutputDir string `yaml:"pipeline_output_dir"`
	} `yaml:"gitlab"`
}

func main() {
	configFile, err := os.ReadFile("../config.yml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	files, err := os.ReadDir(config.TeamCity.DataDir)
	if err != nil {
		log.Fatalf("Error reading input directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".xml") {
			continue
		}

		inputFile := filepath.Join(config.TeamCity.DataDir, file.Name())
		outputFile := filepath.Join(config.GitLab.PipelineOutputDir, strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))+".yml")

		convert.Convert(inputFile, outputFile, config.Templates.JavaTemplate)

		log.Printf("Conversion completed. GitLab pipeline YAML data written to %s", outputFile)
	}
}
