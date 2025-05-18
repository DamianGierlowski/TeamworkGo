package main

import (
	"TeamworkGoTests/customerimporter"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFilePath := flag.String("input", "", "Path to the CSV file")
	outputFilename := flag.String("output", "", "Path to the json file")
	flag.Parse()

	if *csvFilePath == "" {
		log.Fatalf("csv file path is required")
	}

	if *outputFilename == "" {
		log.Fatalf("output json file path is required")
	}

	domainCounts, err := customerimporter.ProcessCSV(*csvFilePath)
	if err != nil {
		log.Fatalf("error processing CSV file '%s': %v", *csvFilePath, err)
	}

	jsonData, err := json.MarshalIndent(domainCounts, "", "  ")
	if err != nil {
		log.Fatalf("error marshalling JSON: %v", err)
	}

	err = saveResponseToFile(*outputFilename, jsonData)
	if err != nil {
		log.Fatalf("error saving data to file '%s': %v", *outputFilename, err)
	}
}

func saveResponseToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", filename, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file '%s': %w", filename, err)
	}

	return nil
}
