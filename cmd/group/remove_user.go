package group

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RemoveUser command
var removeUserCmd = &cobra.Command{
	Use:   "rmuser",
	Short: "Remove a user from a group",
	Long: `Remove a user from an existing group using their email address.
You must provide the group name and the user's email address using the --group and --email flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		group, _ := cmd.Flags().GetString("group")
		email, _ := cmd.Flags().GetString("email")
		if group == "" || email == "" {
			return fmt.Errorf("Both --group and --email flags are required")
		}
		// Placeholder for remove user logic
		fmt.Printf("Removing user %s from group %s\n", email, group)
		return nil
	},
}

func init() {
	removeUserCmd.Flags().StringP("group", "g", "", "Name of the group (required)")
	removeUserCmd.Flags().StringP("email", "e", "", "Email of the user to remove (required)")
	removeUserCmd.MarkFlagRequired("group")
	removeUserCmd.MarkFlagRequired("email")
	GroupCmd.AddCommand(removeUserCmd)
}
