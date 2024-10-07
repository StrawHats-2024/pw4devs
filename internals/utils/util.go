package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ApiResponse represents the structure of the response to be returned
type ApiResponse[T any] struct {
	StatusCode int `json:"statusCode"`
	ResBody    T   `json:"resBody"`
}

const API_URL = "http://159.89.173.5:3000"

// MakeRequest makes an HTTP request with the given parameters and returns an ApiResponse
func MakeRequest[T any](url, method string, body interface{}, authToken string) (*ApiResponse[T], error) {
	// Convert the body to JSON if not nil
	var requestBody []byte
	var err error
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %v", err)
		}
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, API_URL+url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if authToken != "" {
		req.Header.Set("Authorization", "Bearer "+authToken)
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Create a variable of type T for the response body
	var parsedBody T
	err = json.Unmarshal(respBody, &parsedBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// Return the response
	return &ApiResponse[T]{
		StatusCode: resp.StatusCode,
		ResBody:    parsedBody,
	}, nil
}
