/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package secrets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"dv/internals/utils"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a secret",
	Long: `Fetch and display a secret by its ID. You must provide the ID of the secret
to retrieve using the --id or -i flag. By default, this command retrieves the username.
To retrieve the password instead, use the --password or -p flag. 

Additionally, you can use the --copy or -c flag to copy the password to the clipboard,
while still printing the username.`,
	Args: cobra.NoArgs, // No positional arguments, using flags instead
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch the secret ID from the flag
		secretID, _ := cmd.Flags().GetString("id")
		getPassword, _ := cmd.Flags().GetBool("password")
		copyPassword, _ := cmd.Flags().GetBool("copy")

		// Validate that the ID is provided
		if secretID == "" {
			return fmt.Errorf("Secret ID is required. Use --id to specify the ID.")
		}
		id, err := strconv.Atoi(secretID)
		if err != nil {
			return fmt.Errorf("Invalid secret_id")
		}

		res, err := utils.MakeRequest[resBodyGet]("/v1/secrets", http.MethodGet, reqBodyGet{SecretID: id}, utils.GetAuthtoken())
		if err != nil {
			return fmt.Errorf("Failed to make fetch request")
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Request failed with status: %d", res.StatusCode)
		}
		encryptedData := res.ResBody.Data.EncryptedData
		iv := res.ResBody.Data.IV
		decryptedData, err := utils.DecryptAESGCM(string(encryptedData), string(iv), utils.GetEncryptionKey())
		if err != nil {
			return fmt.Errorf("Error while decrypted data")
		}
		credentials, err := utils.ParseJSONToCredentials(decryptedData)
		if err != nil {
			return fmt.Errorf("Config file corrupted")
		}

		if copyPassword {
			fmt.Print(credentials.Username)
			return clipboard.WriteAll(credentials.Password)

		} else if getPassword {
			fmt.Print(credentials.Password)
		} else {
			fmt.Print(credentials.Username)
		}
		return nil
	},
}

type reqBodyGet struct {
	SecretID int `json:"secret_id"`
}
type resBodyGet struct {
	Data    utils.SecretRecord `json:"data"`
	Message string             `json:"message"`
}

// Add flags to the command
func init() {
	getCmd.Flags().StringP("id", "i", "", "ID of the secret to retrieve (required)")
	getCmd.Flags().BoolP("password", "p", false, "get password")
	getCmd.Flags().BoolP("copy", "c", false, "copy password")

	// Mark the id flag as required
	getCmd.MarkFlagRequired("id")

	SecretsCmd.AddCommand(getCmd)
}
