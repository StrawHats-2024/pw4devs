package group

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AddUser command
var addUserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "Add a user to a group",
	Long: `Add a user to an existing group using their email address.
You must provide the group name and the user's email address using the --group and --email flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		group, _ := cmd.Flags().GetString("group")
		email, _ := cmd.Flags().GetString("email")
		if group == "" || email == "" {
			return fmt.Errorf("Both --group and --email flags are required")
		}
		// Placeholder for add user logic
		fmt.Printf("Adding user %s to group %s\n", email, group)
		return nil
	},
}

func init() {
	addUserCmd.Flags().StringP("group", "g", "", "Name of the group (required)")
	addUserCmd.Flags().StringP("email", "e", "", "Email of the user to add (required)")
	addUserCmd.MarkFlagRequired("group")
	addUserCmd.MarkFlagRequired("email")
	GroupCmd.AddCommand(addUserCmd)
}

