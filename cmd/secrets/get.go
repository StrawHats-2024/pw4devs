/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package secrets

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a secret",
	Long: `Fetch and display a secret by its ID. You can provide the ID of the secret
to retrieve as a command flag (--id).`,
	Args: cobra.NoArgs, // No positional arguments, using flags instead
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch the secret ID from the flag
		secretID, _ := cmd.Flags().GetString("id")

		// Validate that the ID is provided
		if secretID == "" {
			return fmt.Errorf("Secret ID is required. Use --id to specify the ID.")
		}

		fmt.Println("secretID: ", secretID)
		return nil
	},
}

// Add flags to the command
func init() {
	getCmd.Flags().StringP("id", "i", "", "ID of the secret to retrieve (required)")

	// Mark the id flag as required
	getCmd.MarkFlagRequired("id")

	SecretsCmd.AddCommand(getCmd)
}
