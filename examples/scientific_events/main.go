package main

import (
	"fmt"
	"github.com/JairoAndresr/collector"
)

const (
	gruplacURL                     = "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000001394"
	scientificEventsDefinitionPath = "examples/scientific_events/scientific_events_definition.yml"
)

func main() {

	api := collector.NewAPI()
	results, err := api.GetContentFromFileDefinition(gruplacURL, scientificEventsDefinitionPath)

	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}
