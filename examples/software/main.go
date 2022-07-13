package main

import (
	"fmt"
	"github.com/JairoAndresR/collector"
)

const (
	gruplacURL             = "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000001394"
	softwareDefinitionPath = "examples/software/software_definition.yml"
)

func main() {

	api := collector.NewAPI()
	results, err := api.GetContentFromFileDefinition(gruplacURL, softwareDefinitionPath)

	if err != nil {
		panic(err)
	}

	fmt.Println(results)

}
