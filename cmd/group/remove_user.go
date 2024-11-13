package group

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"dv/internals/utils"
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

		type reqBody struct {
			GroupName string `json:"group_name"`
			UserEmail string `json:"user_email"`
		}
		type resBody struct {
			Message string `json:"message"`
		}
		res, err := utils.MakeRequest[resBody]("/v1/groups/remove_user", http.MethodPost,
			reqBody{GroupName: group, UserEmail: email}, utils.GetAuthtoken())
		if err != nil {
			return err
		}

		switch res.StatusCode {
		case http.StatusOK:
			return nil
		default:
			fmt.Println("Message: ", res.ResBody)
			return fmt.Errorf("Request failed with status code: %d", res.StatusCode)

		}
	},
}

func init() {
	removeUserCmd.Flags().StringP("group", "g", "", "Name of the group (required)")
	removeUserCmd.Flags().StringP("email", "e", "", "Email of the user to remove (required)")
	removeUserCmd.MarkFlagRequired("group")
	removeUserCmd.MarkFlagRequired("email")
	GroupCmd.AddCommand(removeUserCmd)
}
