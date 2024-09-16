package listmodel

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getSecrets() ([]Secret, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/secrets", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	token, err := ReadFileContent("./token.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Set the cookie
	cookie := &http.Cookie{
		Name:  "token",
		Value: token, // Set your actual token value here
	}
	req.AddCookie(cookie)

	// Make the GET request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	// You can read the response body here if needed
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	secretsResponse := []Secret{}
	err = json.Unmarshal(body, &secretsResponse)
	if err != nil {
		return nil, err
	}

	// Print the response body (or handle it as needed)
	return secretsResponse, nil
}
