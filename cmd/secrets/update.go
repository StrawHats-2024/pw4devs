/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package secrets

import (
	"fmt"

	"github.com/spf13/cobra"
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

		// Ensure that at least one of the update fields is provided
		if name == "" && username == "" && password == "" {
			return fmt.Errorf("At least one of --name, --username, or --password must be provided to update the secret.")
		}

		// No update logic for now, just validating input
		fmt.Println("Update command validated successfully")
		return nil
	},
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
