package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/balle/goSwissNutrientDb/nutrientdb"
	"github.com/balle/goSwissNutrientDb/output"
)

func main() {
	component := flag.String("c", "iron", "component like iron or protein")
	lang := flag.String("l", "en", "language (en, de, it, fr)")
	pdfFilename := flag.String("p", "", "Generate a pdf document from results with given filename")
	jsonFilename := flag.String("j", "", "Generate a json file from results with given filename")
	flag.Parse()

	_, err := nutrientdb.SetLang(*lang)

	if err != nil {
		log.Fatal(err)
	}

	componentId, err := nutrientdb.ComponentId(*component)

	if err != nil || componentId == 0 {
		log.Fatalf("Cannot fetch id for %s: %v", *component, err)
	}

	foodList, err := nutrientdb.GetFoodWithComponent(componentId)

	if err != nil {
		log.Fatalf("Cannot find food with %s: %v", *component, err)
	}

	for _, food := range foodList {
		fmt.Printf("%v\n", food)
	}

	if *pdfFilename != "" {
		title := fmt.Sprintf("Foods for component %s", *component)
		err := output.GeneratePdf(*pdfFilename, title, foodList)

		if err != nil {
			log.Fatalf("Failed generating pdf: %s", err)
		}
	}

	if *jsonFilename != "" {
		err := output.GenerateJson(*jsonFilename, foodList)

		if err != nil {
			log.Fatalf("Failed generating json file: %s", err)
		}
	}
}
