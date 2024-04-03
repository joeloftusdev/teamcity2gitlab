package model

import (
	"encoding/xml"
)

type XMLBuildType struct {
	XMLName  xml.Name    `xml:"build-type"`
	Settings XMLSettings `xml:"settings"`
}

type XMLSettings struct {
	XMLName      xml.Name        `xml:"settings"`
	BuildRunners XMLBuildRunners `xml:"build-runners"`
}
type XMLBuildRunners struct {
	XMLName xml.Name    `xml:"build-runners"`
	Runner  []XMLRunner `xml:"runner"`
}

type XMLParameters struct {
	XMLName xml.Name   `xml:"parameters"`
	Param   []XMLParam `xml:"param"`
}
type XMLParam struct {
	XMLName xml.Name `xml:"param"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type XMLRunner struct {
	XMLName    xml.Name      `xml:"runner"`
	ID         string        `xml:"id,attr"`
	Name       string        `xml:"name,attr"`
	Type       string        `xml:"type,attr"`
	Parameters XMLParameters `xml:"parameters"`
}

type Runners = struct {
	Maven2Runners []Maven2Runner
	SimpleRunners []SimpleRunner
}

type Maven2Runner = struct {
	Goals      string
	RunnerArgs string
	Name       string
	Type       string
}

type SimpleRunner = struct {
	ScriptContent string
	Name          string
	Type          string
}
