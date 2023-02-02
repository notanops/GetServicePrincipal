package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	ga "github.com/sethvargo/go-githubactions"
	"gopkg.in/yaml.v3"
)

func main() {

	// Init Github Action package
	action := ga.New()

	// Get input from action
	stack := action.Getenv("stack-name")

	// Construct the file to read
	filename := "Pulumi." + stack + ".yaml"

	// Read it and store it in a bytes array
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		action.Fatalf("Error reading file %s.\n %s", filename, err)
	}

	// Create a struct to match our yaml (one key, we don't need the others)
	var Pulumi struct {
		Config struct {
			ServicePrincipal string `yaml:"az:servicePrincipalName"`
		} `yaml:"config"`
	}

	// Unmarshall our data into our struct
	err = yaml.Unmarshal(data, &Pulumi)
	if err != nil {
		action.Fatalf("Cannot unmarshall YAML config file : %s\n", err)
	}
	// Get ou value from the struct
	sp := Pulumi.Config.ServicePrincipal

	// Replace and capitalize
	sp = strings.ReplaceAll(sp, "\"", "")
	sp = strings.ReplaceAll(sp, "-", "_")
	sp = strings.ToUpper(sp)

	// Set as an output
	action.SetOutput("service-princpial", sp)

	// For workflow logs
	fmt.Printf("Service principal name is : %s from %s file.", sp, filename)
}
