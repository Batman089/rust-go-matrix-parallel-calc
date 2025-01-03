package tests

import (
	"os"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	folderPath := "./go/generated/testFolder"

	// Create the folder
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create folder: %v", err)
	}

	// Check if the folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		t.Errorf("Expected folder to exist, but it does not")
	}

	// Clean up
	os.RemoveAll(folderPath)
}

func TestCreateFile(t *testing.T) {
	filePath := "./go/generated/testFile.txt"

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	file.Close()

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Expected file to exist, but it does not")
	}

	// Clean up
	os.Remove(filePath)
}
