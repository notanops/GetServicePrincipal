package main

import (
	"fmt"
	"io/ioutil"

	ga "github.com/sethvargo/go-githubactions"
	"gopkg.in/yaml.v2"
)

var data = `
config:
	azure-native:location: FranceCentral
    az:servicePrincipalName: sp-az-it-apim-build
    azure-native:servicePrincipalName: sp-az-it-apim-build
    azure-native:servicePrincipalPassword:
        secure: AAABAB7EFtCctyCyUZhYOoB2QMeo2Pjbw6Qtr6SWPsEdL+Kz3OAhwGoItWVmFENZSWl+dvjumWcfQC9brLKmXijYmt2EOnSm
`

func main() {

	action := ga.New() // Init Github action package
	stack := action.Getenv("stack-name")
	fmt.Printf("Stack name : %s", stack)
	filename := "Pulumi." + stack + ".yaml"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		action.Fatalf("Error reading file.")
	}

	var config map[string]interface{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		action.Fatalf("Can't unmarshall YAML file.")
	}

	fmt.Println(config)

	// servicePrincipal := e.Config.AzServicePrincipalName
	// servicePrincipal = strings.ReplaceAll(servicePrincipal, "-", "_") // Replace - by _
	// servicePrincipal = strings.ToUpper(servicePrincipal)              // Capitalize the string
	// action.SetOutput("service-principal", servicePrincipal)           // Set as action output
	// os.Setenv("GITHUB_OUTPUT", servicePrincipal)                      // Set as $GITHUB_OUTPUT
}
