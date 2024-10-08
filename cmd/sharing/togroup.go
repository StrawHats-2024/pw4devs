package sharing

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ToGroup command
var toGroupCmd = &cobra.Command{
	Use:   "togroup",
	Short: "Share a secret with a group",
	Long: `Grant access to a secret for all members of a specified group.
You must provide the ID of the secret and the name of the group using the --secret-id and --group flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		secretID, _ := cmd.Flags().GetString("secret-id")
		group, _ := cmd.Flags().GetString("group")
		if secretID == "" || group == "" {
			return fmt.Errorf("Both --secret-id and --group flags are required")
		}
		// Placeholder for share to group logic
		fmt.Printf("Sharing secret %s with group %s\n", secretID, group)
		return nil
	},
}

func init() {
	toGroupCmd.Flags().StringP("secret-id", "s", "", "ID of the secret to share (required)")
	toGroupCmd.Flags().StringP("group", "g", "", "Name of the group to share with (required)")
	toGroupCmd.MarkFlagRequired("secret-id")
	toGroupCmd.MarkFlagRequired("group")
	ShareCmd.AddCommand(toGroupCmd)
}
