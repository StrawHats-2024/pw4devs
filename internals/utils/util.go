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

// const API_URL = "http://159.89.173.5:3000"

const API_URL = "http://localhost:4000"

// MakeRequest makes an HTTP request with the given parameters and returns an ApiResponse
func MakeRequest[T any](url, method string, body interface{}, authToken string) (*ApiResponse[T], error) {
	// Convert the body to JSON if provided and add it to the request
	var requestBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %v", err)
		}
		requestBody = bytes.NewBuffer(b)
	}

	// Create and configure the HTTP request
	req, err := http.NewRequest(method, API_URL+url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if authToken != "" {
		req.Header.Set("Authorization", "Bearer "+authToken)
	}

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response has a body
	if resp.ContentLength == 0 {
		return &ApiResponse[T]{StatusCode: resp.StatusCode}, nil
	}

	// Parse the response body if present
	var parsedBody T
	if err := json.NewDecoder(resp.Body).Decode(&parsedBody); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// Return the response
	return &ApiResponse[T]{StatusCode: resp.StatusCode, ResBody: parsedBody}, nil
}
