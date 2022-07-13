package main

import (
	"fmt"
	"github.com/JairoAndresR/collector"
)

const (
	gruplacURL          = "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000001394"
	booksDefinitionPath = "examples/books/books_definition.yml"
)

func main() {

	api := collector.NewAPI()
	results, err := api.GetContentFromFileDefinition(gruplacURL, booksDefinitionPath)

	if err != nil {
		panic(err)
	}

	fmt.Println(results)

}
