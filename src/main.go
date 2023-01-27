package main

import (
	"os"
	"strings"

	ga "github.com/sethvargo/go-githubactions"
	"gopkg.in/yaml.v2"
)

var data = `
config:
    az:servicePrincipalName: sp-az-it-apim-build
    azure-native:servicePrincipalName: sp-az-it-apim-build
    azure-native:servicePrincipalPassword:
        secure: AAABAB7EFtCctyCyUZhYOoB2QMeo2Pjbw6Qtr6SWPsEdL+Kz3OAhwGoItWVmFENZSWl+dvjumWcfQC9brLKmXijYmt2EOnSm
    azure-native:location: FranceCentral
`

// PulumiConf represent Pulumi config file
type PulumiConf struct {
	Config struct {
		AzServicePrincipalName string `yaml:"az:servicePrincipalName"`
	} `yaml:"config"`
}

func main() {

	action := ga.New() // Init Github action package
	e := PulumiConf{}
	err := yaml.Unmarshal([]byte(data), &e)
	if err != nil {
		action.Fatalf("Can't unmarshall Pulumi config file. Does the yaml key/value az:servicePrincipalName is set ?")
	}
	servicePrincipal := e.Config.AzServicePrincipalName
	servicePrincipal = strings.ReplaceAll(servicePrincipal, "-", "_") // Replace - by _
	servicePrincipal = strings.ToUpper(servicePrincipal)              // Capitalize the string
	action.SetOutput("service-principal", servicePrincipal)           // Set as action output
	os.Setenv("GITHUB_OUTPUT", servicePrincipal)                      // Set as $GITHUB_OUTPUT
}
