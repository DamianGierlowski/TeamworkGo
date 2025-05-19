package app

import (
	"TeamworkGoTests/internal/customerimporter"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ProcessAndSave(inputFilePath string, outputFilename string) error {
	domainCounts, err := customerimporter.ProcessCSV(inputFilePath)
	if err != nil {
		return fmt.Errorf("error processing CSV file '%s': %w", inputFilePath, err)
	}

	jsonData, err := json.MarshalIndent(domainCounts, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	err = saveResponseToFile(outputFilename, jsonData)
	if err != nil {
		return fmt.Errorf("error saving data to file '%s': %w", outputFilename, err)
	}

	return nil
}

func saveResponseToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", filename, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("Error closing file %s: %v", filename, closeErr)
		}
	}()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file '%s': %w", filename, err)
	}

	return nil
}
