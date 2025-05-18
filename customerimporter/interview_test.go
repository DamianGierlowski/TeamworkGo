package customerimporter_test

import (
	"TeamworkGoTests/customerimporter"
	"os"

	"reflect"
	"sort"
	"testing"
)

func sortDomainCounts(counts []customerimporter.DomainCount) {
	sort.Slice(counts, func(i, j int) bool {

		return counts[i].Domain < counts[j].Domain

	})
}

// Define sample data for testing
const sampleCSVData = `first_name,last_name,email,gender,ip_address
Mildred,Hernandez,mhernandez0@github.io,Female,38.194.51.128
Bonnie,Ortiz,bortiz1@cyberchimps.com,Female,197.54.209.129
Dennis,Henry,dhenry2@github.io,Male,155.75.186.217
Justin,Hansen,jhansen3@360.cn,Male,251.166.224.119
Rose,Watson,rwatson1w@shutterfly.com,Female,39.148.64.108
Dennis,Wilson,dwilson1x@github.io,Male,118.191.194.30
Maria,Burke,mburke1y@shutterfly.com,Female,25.58.184.182
Denise,Ford,dford1z@issuu.com,Female,184.37.246.43
Brenda,Gonzales,bgonzales20@shutterfly.com,Female,19.244.234.41
`

func TestProcessCSV_EmptyFile(t *testing.T) {

	tmpFilePath := createTempFile(t, "")

	counts, err := customerimporter.ProcessCSV(tmpFilePath)
	if err != nil {
		t.Fatalf("Expected no error for empty file, got: %v", err)
	}

	if len(counts) != 0 {
		t.Errorf("Expected empty domain counts for empty file, got: %v", counts)
	}
}

func TestProcessCSV_NonExistentFile(t *testing.T) {

	_, err := customerimporter.ProcessCSV("non_existent_file.csv")
	if err == nil {
		t.Fatal("Expected an error for non-existent file, got nil")
	}
}

func TestProcessCSV_SampleData(t *testing.T) {
	tmpFilePath := createTempFile(t, sampleCSVData)
	defer os.Remove(tmpFilePath) // Clean up the temporary file after the test
	// Process the temporary file
	counts, err := customerimporter.ProcessCSV(tmpFilePath)
	if err != nil {
		t.Fatalf("Expected no error for sample data, got: %v", err)
	}

	// Define expected counts
	expectedCounts := []customerimporter.DomainCount{
		{Domain: "github.io", Count: 3},
		{Domain: "cyberchimps.com", Count: 1},
		{Domain: "360.cn", Count: 1},
		{Domain: "shutterfly.com", Count: 3},
		{Domain: "issuu.com", Count: 1},
	}

	sortDomainCounts(counts)
	sortDomainCounts(expectedCounts)

	if !reflect.DeepEqual(counts, expectedCounts) {
		t.Errorf("Expected domain counts %v, got %v", expectedCounts, counts)
	}
}

func createTempFile(t *testing.T, data string) string {
	// Create a temporary file. The second argument is the pattern for the filename.
	tmpFile, err := os.CreateTemp("", "sample_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	// Get the path before deferring close and remove
	tmpFilePath := tmpFile.Name()
	// Write the data to the temporary file
	_, err = tmpFile.Write([]byte(data))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// We must close the file before ProcessCSV tries to open it again
	err = tmpFile.Close()
	if err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpFilePath
}
