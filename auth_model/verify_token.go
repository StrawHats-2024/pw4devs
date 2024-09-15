package authmodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type VerifyTokenResponse struct {
	Valid   bool   `json:"valid"`
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

func VerifyToken(token string) (*VerifyTokenResponse, error) {
	url := "http://localhost:3000/api/v1/auth/verify" // Update with the actual API URL

	// Create request payload
	requestBody, err := json.Marshal(VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Make POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to make post request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Handle non-200 response status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var verifyResp VerifyTokenResponse
	err = json.Unmarshal(body, &verifyResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &verifyResp, nil
}
