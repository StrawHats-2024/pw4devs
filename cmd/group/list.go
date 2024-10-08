package group

import (
	"fmt"

	"github.com/spf13/cobra"
)

// List command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all groups",
	Long: `Display a list of all groups accessible to the current user.
You can specify the maximum number of groups to display and the number of groups to skip using the --limit and --offset flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")
		// Placeholder for list logic
		fmt.Printf("Listing groups with limit %d and offset %d\n", limit, offset)
		return nil
	},
}

func init() {
	listCmd.Flags().IntP("limit", "l", 10, "Maximum number of groups to display (default: 10)")
	listCmd.Flags().IntP("offset", "o", 0, "Number of groups to skip (for pagination)")
	GroupCmd.AddCommand(listCmd)
}
