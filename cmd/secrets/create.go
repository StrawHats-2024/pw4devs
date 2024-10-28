package secrets

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new secret",
	Long: `Add a new secret with a name, username, password, and type. You can provide the name, username, password,
and type as command flags. If no flags are provided, an interactive form will be presented for input.`,
	Args: cobra.NoArgs, // No positional arguments, using flags instead
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, username, password string

		// Fetch flag values
		name, _ = cmd.Flags().GetString("name")
		username, _ = cmd.Flags().GetString("username")
		password, _ = cmd.Flags().GetString("password")

		// Check if all flags are provided, else prompt for interactive input
		if name == "" || username == "" || password == "" {
			// Interactive TUI form for input
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name").Value(&name),
					huh.NewInput().Title("Username").Value(&username),
					huh.NewInput().Title("Password").EchoMode(huh.EchoModePassword).Value(&password),
				),
			)
			form.Run()
		}

		// Check if all required values are filled
		if name == "" || username == "" || password == "" {
			return fmt.Errorf("All fields (name, username, password) are required.")
		}
    return utils.CreateSecret(name, username, password)
	},
}

// Add flags to the command
func init() {
	createCmd.Flags().StringP("name", "n", "", "Name of the secret (required)")
	createCmd.Flags().StringP("username", "u", "", "Username for the secret (required)")
	createCmd.Flags().StringP("password", "p", "", "Password for the secret (required)")

	SecretsCmd.AddCommand(createCmd)
}
