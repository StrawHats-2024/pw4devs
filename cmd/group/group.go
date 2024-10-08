package group

import (
	"github.com/spf13/cobra"
)

// SecretsCmd represents the secrets command
var GroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage user groups",
	Long: `The group command allows you to manage user groups within the password manager.
You can create, retrieve, update, delete, and manage users within groups using this command.`,
}
