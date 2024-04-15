# TeamCity to GitLab migration utility

## Prerequisites
* Golang v1.21.6 or later

## Overview
Simple utility that converts a TeamCity build configuration XML document into a GitLab pipeline .yml file.

It loads a set of user defined templates from a location defined in a configuration file `config.yml` & uses them to transform between the two formats.
## Design

![design](https://github.com/joeloftusdev/teamcity2gitlab/assets/152509645/7bc67d06-9e65-4539-b46c-b0b91fedef55)

## Build & Run
 Rename `config.yml.default` to `config.yml` & specify the location of your TeamCity config data directory plus your output directory for GitLab pipelines.

 If you leave the default values as is the utility will use the example TeamCity config files in the xml directory.

To run the app:


````
cd cmd
go run .
````


## Info
Currently this app supports a simple build configuration using Maven & SimpleRunner build steps. 

I'll be adding templates for other languages & build tools in the future.

## Experimental

I have created a test environment/lab containing a Gitlab server & runner plus a TeamCity server & agent. 

You can read more about this [here](/lab/lab.md).

You will need a Docker environment set up running locally including Docker compose.


