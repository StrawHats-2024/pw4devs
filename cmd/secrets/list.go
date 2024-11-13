/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package secrets

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"dv/internals/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secrets",
	Long: `Display a list of all secrets accessible to the current user.
You can control the number of secrets displayed using the --limit and --offset flags for pagination.`,
	Args: cobra.NoArgs, // No positional arguments, using flags instead
	RunE: func(cmd *cobra.Command, args []string) error {

		type resBody struct {
			Data    []utils.SecretRecord `json:"data"`
			Message string               `json:"message"`
		}
		// Here we would normally call the logic to fetch and list secrets.
		// Currently, just validating inputs and placeholder message.
		res, err := utils.MakeRequest[resBody]("/v1/secrets/user", http.MethodGet, nil, utils.GetAuthtoken())
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Fetch request failed with status code: %d", res.StatusCode)
		}
		data := res.ResBody.Data
		if len(data) > 0 {
			printSecrets(data)
		}

		return nil
	},
}

// Add flags to the command
func init() {

	SecretsCmd.AddCommand(listCmd)
}
func printSecrets(secrets []utils.SecretRecord) {
	// Print a plain-text table header of secrets for piping to `fzf`
	fmt.Println("ID\tName\t\tCreated At")
	fmt.Println("---------------------------------------------------")

	// Format each secret's details
	for _, secret := range secrets {
		// Format `CreatedAt` to a more readable format
		formattedTime := secret.CreatedAt.Format("Jan 02, 2006 03:04 PM")
		fmt.Printf("%d\t%s\t\t%s\n", secret.ID, secret.Name, formattedTime)
	}
}
