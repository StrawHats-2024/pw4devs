package group

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve group information",
	Long: `Display details of a group by its name.
You must provide the name of the group using the --name flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			return fmt.Errorf("The --name flag is required")
		}
		// Placeholder for get logic
		fmt.Printf("Getting group information for: %s\n", name)
		return nil
	},
}

func init() {
	getCmd.Flags().StringP("name", "n", "", "Name of the group to retrieve (required)")
	getCmd.MarkFlagRequired("name")
	GroupCmd.AddCommand(getCmd)
}
