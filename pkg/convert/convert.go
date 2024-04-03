package convert

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"teamcity2gitlab/pkg/model"
	"text/template"
)

func Convert(inputFile, outputFile, templateFile string) {

	xmlFile, err := os.Open(inputFile)
	if err != nil {
		log.Printf("Error opening XML file %s: %v", inputFile, err)
		return
	}

	defer xmlFile.Close()

	xmlBytes, err := io.ReadAll(xmlFile)
	if err != nil {
		log.Printf("Error reading XML file %s: %v", inputFile, err)
		return
	}

	if len(xmlBytes) == 0 {
		log.Printf("Empty XML file: %s", inputFile)
		return
	}

	var xmlData model.XMLBuildType
	err = xml.Unmarshal(xmlBytes, &xmlData)
	if err != nil {
		log.Printf("Error parsing XML file %s: %v", inputFile, err)
		return
	}

	runners := xmlData.Settings.BuildRunners.Runner

	myRunners := BuildTemplateModel(runners)

	tmpl, err := template.ParseFiles(templateFile) //take file from config
	if err != nil {
		log.Fatalf("Errors loading template: %s", err)
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Printf("Error creating output file %s: %v", outputFile, err)
		return
	}
	defer outFile.Close()

	err = tmpl.Execute(outFile, myRunners)
	if err != nil {
		log.Printf("Error executing YAML template: %v", err)
		return
	}

	log.Printf("Conversion completed. YAML data written to %s", outputFile)

}

func BuildTemplateModel(runners []model.XMLRunner) model.Runners {

	maven2Runners := []model.Maven2Runner{}
	simpleRunners := []model.SimpleRunner{}

	for _, runner := range runners {
		if runner.Type == "Maven2" {
			maven2Runner := model.Maven2Runner{
				Type: "Maven2",
				Name: runner.Name,
			}
			params := runner.Parameters.Param
			for _, param := range params {
				switch param.Name {
				case "goals":
					maven2Runner.Goals = param.Value
				case "runnerArgs":
					maven2Runner.RunnerArgs = param.Value
				}
			}
			maven2Runners = append(maven2Runners, maven2Runner)
		} else if runner.Type == "simpleRunner" {
			simpleRunner := model.SimpleRunner{
				Type: "simpleRunner",
				Name: runner.Name,
			}
			params := runner.Parameters.Param
			for _, param := range params {
				switch param.Name {
				case "script.content":
					simpleRunner.ScriptContent = param.Value
				}
			}
			simpleRunners = append(simpleRunners, simpleRunner)
		}
	}
	return model.Runners{
		Maven2Runners: maven2Runners,
		SimpleRunners: simpleRunners,
	}
}
