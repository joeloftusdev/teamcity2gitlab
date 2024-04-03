package convert_test

import (
	"teamcity2gitlab/pkg/convert"
	"testing"

	"teamcity2gitlab/pkg/model"

	"github.com/stretchr/testify/assert"
)

func Test_BuildTemplateModel(t *testing.T) {
	xmlRunners := []model.XMLRunner{
		{
			Name: "Test1",
			Type: "Maven2",
			Parameters: model.XMLParameters{
				Param: []model.XMLParam{
					{Name: "goals", Value: "clean install"},
					{Name: "runnerArgs", Value: "-DskipTests"},
				},
			},
		},
		{
			Name: "Test2",
			Type: "simpleRunner",
			Parameters: model.XMLParameters{
				Param: []model.XMLParam{
					{Name: "script.content", Value: "echo 'Hello, world!'"},
				},
			},
		},
	}

	myRunners := convert.BuildTemplateModel(xmlRunners)

	assert.Equal(t, len(myRunners.Maven2Runners), 1, "Expected 1 Maven2Runner")
	assert.Equal(t, len(myRunners.SimpleRunners), 1, "Expected 1 SimpleRunner")
	assert.Equal(t, myRunners.Maven2Runners[0].Name, "Test1", "Expected name: Test1")
	assert.Equal(t, myRunners.SimpleRunners[0].Type, "simpleRunner", "Expected type: simpleRunner")

}
