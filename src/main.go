package main

import (
	"fmt"
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

// PulumiConf represent Pulumi config
type PulumiConf struct {
	Config struct {
		AzServicePrincipalName string `yaml:"az:servicePrincipalName"`
	} `yaml:"config"`
}

func main() {

	action := ga.New()
	e := PulumiConf{}
	yaml.Unmarshal([]byte(data), &e)
	servicePrincipal := e.Config.AzServicePrincipalName
	servicePrincipal = strings.ReplaceAll(servicePrincipal, "-", "_")
	servicePrincipal = strings.ToUpper(servicePrincipal)
	fmt.Println(servicePrincipal)
	action.SetOutput("service-principal", servicePrincipal)
}
