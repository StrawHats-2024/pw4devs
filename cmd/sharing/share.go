package sharing

import "github.com/spf13/cobra"

var ShareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share secrets with users or groups",
	Long: `The share command allows you to share secrets with specific users or groups.
You can grant access to a secret for all members of a specified group or a specific user using their email address.`,
}
