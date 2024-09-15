package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("file not found")

func readFileContent(filePath string) (string, error) {
	// Read the entire content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		// Check if the error is because the file doesn't exist
		if os.IsNotExist(err) {
			return "", ErrFileNotFound
		}
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Convert the content to a string and return
	return string(content), nil
}
