package main

import (
	"TeamworkGoTests/customerimporter"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	csvFilePath := "customers.csv"

	domainCounts, err := customerimporter.ProcessCSV(csvFilePath)
	if err != nil {
		log.Fatalf("Error processing CSV file '%s': %v", csvFilePath, err)
	}

	fmt.Println("Successfully processed CSV. Domain Counts:")
	jsonData, err := json.MarshalIndent(domainCounts, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}
