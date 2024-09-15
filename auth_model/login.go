package authmodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// AuthRequest represents the request structure for the login API.
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the response structure from the login API.
type AuthResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

// Login function to authenticate the user and store the JWT token locally.
func Login(email, password string) error {
	const url = "http://localhost:3000/api/v1/auth/login" // Change to your API endpoint

	// Create the request payload
	authRequest := AuthRequest{
		Email:    email,
		Password: password,
	}

	// Convert the payload to JSON
	payload, err := json.Marshal(authRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	// Make the POST request
	res, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	// Check the response status
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %d: %s", res.StatusCode, body)
	}

	// Unmarshal the response to get the token
	var authResponse AuthResponse
	if err := json.Unmarshal(body, &authResponse); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	// Store the token in a local file
	if err := os.WriteFile("token.txt", []byte(authResponse.Token), 0644); err != nil {
		return fmt.Errorf("failed to write token to file: %v", err)
	}

	fmt.Println("Token stored in token.txt")
	return nil
}
