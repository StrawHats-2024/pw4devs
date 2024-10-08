package sharing

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ToUser command
var toUserCmd = &cobra.Command{
	Use:   "touser",
	Short: "Share a secret with a user",
	Long: `Grant access to a secret for a specific user by their email address.
You must provide the ID of the secret and the user's email address using the --secret-id and --email flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		secretID, _ := cmd.Flags().GetString("secret-id")
		email, _ := cmd.Flags().GetString("email")
		if secretID == "" || email == "" {
			return fmt.Errorf("Both --secret-id and --email flags are required")
		}
		// Placeholder for share to user logic
		fmt.Printf("Sharing secret %s with user %s\n", secretID, email)
		return nil
	},
}

func init() {
	toUserCmd.Flags().StringP("secret-id", "s", "", "ID of the secret to share (required)")
	toUserCmd.Flags().StringP("email", "e", "", "Email of the user to share with (required)")
	toUserCmd.MarkFlagRequired("secret-id")
	toUserCmd.MarkFlagRequired("email")
	ShareCmd.AddCommand(toUserCmd)
}
