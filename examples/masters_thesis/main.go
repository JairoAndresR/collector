package main

import (
	"fmt"
	"github.com/JairoAndresR/collector"
)

const (
	gruplacURL                 = "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000001394"
	managedWorksDefinitionPath = "examples/masters_thesis/masters_thesis_definition.yml"
)

func main() {

	api := collector.NewAPI()
	results, err := api.GetContentFromFileDefinition(gruplacURL, managedWorksDefinitionPath)

	if err != nil {
		panic(err)
	}

	for _, fields := range results {
		for field, value := range fields {
			fmt.Printf("%s: %s\n", field, value)
		}
		fmt.Println()
	}
}
