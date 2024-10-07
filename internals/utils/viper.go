package utils

import (
	"os"

	"github.com/spf13/viper"
)

const AuthTokenKey = "auth_token"
const UserEmailKey = "user_email"

func SetAuthToken(token string) error {
	// Set the auth token in Viper
	viper.Set(AuthTokenKey, token)
	return writeConfig()
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
