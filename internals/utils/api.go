package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateSecret(name, username, password string) error {
	data := map[string]string{
		"username": username,
		"password": password,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	plaintext := string(jsonData)
	encryptedData, iv, err := EncryptAESGCM(plaintext, GetEncryptionKey())
  fmt.Println("iv: ", iv);
  fmt.Println("encryptedData: ", encryptedData);
	if err != nil {
		return err
	}
	type reqBody struct {
		EncryptedData string `json:"encrypted_data"`
		IV            string `json:"iv"`
		Name          string `json:"name"`
	}

	type resBody struct {
		Message  string `json:"message"`
		SecretID int    `json:"secret_id"`
	}
	res, err := MakeRequest[resBody]("/v1/secrets",
		http.MethodPost,
		reqBody{EncryptedData: encryptedData, IV: iv, Name: name},
		GetAuthtoken())

	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusCreated {
		fmt.Printf("Secret created successfully: %v", res.ResBody.SecretID)
	} else {
		fmt.Printf("Error StatusCode: %d", res.StatusCode)
	}
	return nil
}
