package group

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a group",
	Long: `Remove a group from the system by its name.
You must provide the name of the group using the --name flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			return fmt.Errorf("The --name flag is required")
		}
		// Placeholder for delete logic
		fmt.Printf("Deleting group: %s\n", name)
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringP("name", "n", "", "Name of the group to delete (required)")
	deleteCmd.MarkFlagRequired("name")
	GroupCmd.AddCommand(deleteCmd)
}
