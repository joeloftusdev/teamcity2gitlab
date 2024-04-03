# TeamCity to GitLab migration utility

## Prerequisites
* Golang v1.21.6 or later

## Overview
Simple utility that converts a TeamCity build configuration XML document into a GitLab pipeline .yml file.

It loads a set of user defined templates from a location defined in a configuration file `config.yml` & uses them to transform between the two formats.
## Build & Run
 Rename `config.yml.default` to `config.yml` & specify the location of your TeamCity config data directory plus your output directory for GitLab pipelines.

 If you leave the default values as is the utility will use the example TeamCity config files in the xml directory.

To run the app:


````
cd cmd/app
go run .
````

## Info
Currently this app supports a simple build configuration using Maven & SimpleRunner build steps. 

I'll be adding templates for other languages & build tools in the future.
