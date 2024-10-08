package group

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a group's name",
	Long: `Change the name of an existing group.
You must provide both the current name and the new name using the --oldname and --newname flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		oldname, _ := cmd.Flags().GetString("oldname")
		newname, _ := cmd.Flags().GetString("newname")
		if oldname == "" || newname == "" {
			return fmt.Errorf("Both --oldname and --newname flags are required")
		}
		// Placeholder for update logic
		fmt.Printf("Updating group from %s to %s\n", oldname, newname)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("oldname", "o", "", "Current name of the group (required)")
	updateCmd.Flags().StringP("newname", "n", "", "New name for the group (required)")
	updateCmd.MarkFlagRequired("oldname")
	updateCmd.MarkFlagRequired("newname")
	GroupCmd.AddCommand(updateCmd)
}
