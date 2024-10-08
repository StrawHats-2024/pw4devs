/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package secrets

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a secret",
	Long: `Remove a secret from the system by its ID.
You must provide the ID of the secret to be deleted using the --id flag.`,
	Args: cobra.NoArgs, // No positional arguments, using flags instead
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch the 'id' flag
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			return fmt.Errorf("The --id flag is required")
		}

		// Placeholder for delete logic
		fmt.Printf("Deleting secret with ID: %s\n", id)

		return nil
	},
}

// Add flags to the command
func init() {
	deleteCmd.Flags().StringP("id", "i", "", "ID of the secret to delete (required)")

	// Mark 'id' as a required flag
	deleteCmd.MarkFlagRequired("id")

	SecretsCmd.AddCommand(deleteCmd)
}
