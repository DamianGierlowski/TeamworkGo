// Package customerimporter reads from a CSV file and returns a sorted (data
// structure of your choice) of email domains along with the number of customers
// with e-mail addresses for each domain. This should be able to be ran from the
// CLI and output the sorted domains to the terminal or to a file. Any errors
// should be logged (or handled). Performance matters (this is only ~3k lines,
// but could be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type DomainCount struct {
	Domain string `json:"domain"`
	Count  int    `json:"count"`
}

func ProcessCSV(filePath string) ([]DomainCount, error) {

	domainCounts := make(map[string]int)

	reader, closeFunc, err := readFromFile(filePath)

	if err != nil {

		return nil, err
	}

	defer closeFunc()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %w", err)
		}

		if len(record) > 2 {
			email := record[2]
			parts := strings.Split(email, "@")
			if len(parts) == 2 && verifyDomain(parts[1]) {
				domain := parts[1]
				domainCounts[domain]++
			} else {
				log.Printf("invalid email format: %s", email)
			}
		}
	}

	var domainCountList []DomainCount
	for domain, count := range domainCounts {
		domainCountList = append(domainCountList, DomainCount{Domain: domain, Count: count})
	}

	return domainCountList, nil
}

func readFromFile(filePath string) (*csv.Reader, func() error, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("opening file: %w", err)
	}
	closeFunc := func() error {
		return file.Close()
	}

	reader := csv.NewReader(file)

	return reader, closeFunc, nil
}

func verifyDomain(domain string) bool {
	return strings.Contains(domain, ".")
}
