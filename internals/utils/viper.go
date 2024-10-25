package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const AuthTokenKey = "auth_token"
const UserEmailKey = "user_email"
const EncryptionKey = "key"

func SetAuthToken(token string) error {
	// Set the auth token in Viper
	viper.Set(AuthTokenKey, token)
	return writeConfig()
}

func GetAuthtoken() string {
	token := viper.GetString(AuthTokenKey)
	if token == "" {
		fmt.Println("Auth Token not found. Please login first")
		os.Exit(1)
		return ""
	}
	return token
}
func GetEncryptionKey() string {
	key := viper.GetString(EncryptionKey)
	if key == "" {
		fmt.Println("EncryptionKey not found.")
		fmt.Println("you can set your key using: auth set-key")
		os.Exit(1)
		return ""
	}
	return key
}

func SetUserEmail(email string) error {
	viper.Set(UserEmailKey, email)
	return writeConfig()
}

func writeConfig() error {

	// Write or overwrite the configuration file
	if err := viper.WriteConfig(); err != nil {
		// If the config file doesn't exist yet, create a new one
		if os.IsNotExist(err) {
			if err := viper.SafeWriteConfig(); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
