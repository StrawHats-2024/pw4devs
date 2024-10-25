package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

// Helper function to convert Base64 to byte slice
func base64ToBytes(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}

func DecryptAESGCM(ciphertextBase64, ivBase64, keyBase64 string) (string, error) {
	// Step 1: Decode the Base64-encoded key, IV, and ciphertext
	key, err := base64ToBytes(keyBase64)
	if err != nil {
		return "", err
	}

	iv, err := base64ToBytes(ivBase64)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64ToBytes(ciphertextBase64)
	if err != nil {
		return "", err
	}

	// Step 2: Import the AES key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Step 3: Create the GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Step 4: Decrypt the data using AES-GCM
	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}

	// Step 5: Convert decrypted byte slice to string and return
	return string(plaintext), nil
}

// Helper function to convert string to byte slice
func stringToBytes(s string) []byte {
	return []byte(s)
}

// Helper function to convert byte slice to Base64 string
func bytesToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func EncryptAESGCM(plaintext, keyBase64 string) (string, string, error) {
	// Step 1: Decode the Base64-encoded key
	key, err := base64ToBytes(keyBase64)
	if err != nil {
		return "", "", err
	}

	// Step 2: Generate a random IV (Initialization Vector)
	iv := make([]byte, 12) // AES-GCM requires a 12-byte IV
	if _, err := rand.Read(iv); err != nil {
		return "", "", err
	}

	// Step 3: Import the AES key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	// Step 4: Create the GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	// Step 5: Encrypt the plaintext using AES-GCM
	ciphertext := gcm.Seal(nil, iv, stringToBytes(plaintext), nil)

	// Step 6: Return the ciphertext and IV as Base64 strings
	return bytesToBase64(ciphertext), bytesToBase64(iv), nil
}
