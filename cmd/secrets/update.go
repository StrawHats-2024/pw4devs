/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package secrets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"dv/internals/utils"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing secret",
	Long: `Modify one or more fields of an existing secret by providing its ID.
You can change the secret's name, username, and/or password by specifying the corresponding flags.`,
	Args: cobra.NoArgs, // No positional arguments, using flags instead
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch flag values
		secretID, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		// Validate that the ID is provided
		if secretID == "" {
			return fmt.Errorf("Secret ID is required. Use --id to specify the ID.")
		}

		id, err := strconv.Atoi(secretID)
		if err != nil {
			return fmt.Errorf("Invalid secret_id")
		}

		//TODO: impliment partial update
		if name == "" || username == "" || password == "" {
			return fmt.Errorf("All --name, --username, or --password must be provided to update the secret.")
		}
		data := map[string]string{
			"username": username,
			"password": password,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}

		plaintext := string(jsonData)
		encryptedData, iv, err := utils.EncryptAESGCM(plaintext, utils.GetEncryptionKey())
		if err != nil {
			return err
		}
		res, err := utils.MakeRequest[any]("/v1/secrets", http.MethodPatch, reqbodyUpdate{SecretID: id,
			EncryptedData: encryptedData,
			IV:            iv,
			Name:          name,
		}, utils.GetAuthtoken())
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Request failed with code %d", res.StatusCode)
		}
		return nil
	},
}

type reqbodyUpdate struct {
	EncryptedData string `json:"encrypted_data"`
	IV            string `json:"iv"`
	Name          string `json:"name"`
	SecretID      int    `json:"secret_id"`
}

// Add flags to the command
func init() {
	updateCmd.Flags().StringP("id", "i", "", "ID of the secret to update (required)")
	updateCmd.Flags().String("name", "", "New name for the secret")
	updateCmd.Flags().String("username", "", "New username for the secret")
	updateCmd.Flags().String("password", "", "New password for the secret")

	// Mark the id flag as required
	updateCmd.MarkFlagRequired("id")

	SecretsCmd.AddCommand(updateCmd)
}
