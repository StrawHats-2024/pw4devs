/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package secrets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a secret",
	Long: `Remove a secret from the system by its ID.
You must provide the ID of the secret to be deleted using the --id flag.`,
	Args: cobra.NoArgs, // No positional arguments, using flags instead
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch the 'secretID' flag
		secretID, _ := cmd.Flags().GetString("id")
		if secretID == "" {
			return fmt.Errorf("The --id flag is required")
		}
		id, err := strconv.Atoi(secretID)
		if err != nil {
			return err
		}

		res, err := utils.MakeRequest[any]("/v1/secrets", http.MethodDelete, reqBodyDel{id}, utils.GetAuthtoken())
		if err != nil {
			return err
		}
		if res.StatusCode == http.StatusNoContent {
			return nil
		}
		switch res.StatusCode {
		case http.StatusUnauthorized:
			fmt.Print("Not authorised for this action")
		case http.StatusUnprocessableEntity:
			fmt.Print("Invalid request")
		case http.StatusNotFound:
			fmt.Printf("\nNo secret found with id: %s", secretID)
		}
		return nil
	},
}

type reqBodyDel struct {
	SecretID int `json:"secret_id"`
}

// Add flags to the command
func init() {
	deleteCmd.Flags().StringP("id", "i", "", "ID of the secret to delete (required)")

	// Mark 'id' as a required flag
	deleteCmd.MarkFlagRequired("id")

	SecretsCmd.AddCommand(deleteCmd)
}
