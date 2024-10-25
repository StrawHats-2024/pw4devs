package utils

import "encoding/json"

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Function to convert JSON string to struct
func ParseJSONToCredentials(jsonStr string) (*Credentials, error) {
	// Create an instance of the struct to hold the parsed data
	var creds Credentials

	// Parse the JSON string into the struct
	err := json.Unmarshal([]byte(jsonStr), &creds)
	if err != nil {
		return nil, err
	}

	// Return the parsed struct and no error
	return &creds, nil
}
