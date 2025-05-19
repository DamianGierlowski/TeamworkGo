package main

import (
	"TeamworkGoTests/internal/app"
	"flag"
	"log"
)

func main() {
	inputFilePath := flag.String("input", "", "Path to the CSV file")
	outputFilename := flag.String("output", "", "Path to the json file")
	flag.Parse()

	if *inputFilePath == "" {
		log.Fatalf("csv file path is required")
	}

	if *outputFilename == "" {
		log.Fatalf("output json file path is required")
	}

	if err := app.ProcessAndSave(*inputFilePath, *outputFilename); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
