package group

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new group",
	Long: `Create a new group with a specified name.
You must provide the name of the new group using the --name flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			return fmt.Errorf("The --name flag is required")
		}
		// Placeholder for create logic
		fmt.Printf("Creating group: %s\n", name)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "Name of the new group (required)")
	createCmd.MarkFlagRequired("name")
	GroupCmd.AddCommand(createCmd)
}
