package group

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"dv/internals/utils"
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
		type reqBody struct {
			GroupName string `json:"group_name"`
			UserEmail string `json:"user_email"`
		}
		type resBody struct {
			Message string `json:"message"`
		}
		res, err := utils.MakeRequest[resBody]("/v1/groups/add_user", http.MethodPost,
			reqBody{GroupName: group, UserEmail: email}, utils.GetAuthtoken())
		if err != nil {
			return err
		}

		switch res.StatusCode {
		case http.StatusOK:
			return nil
		default:
			fmt.Println("Message: ", res.ResBody.Message)
			return fmt.Errorf("Request failed with status code: %d", res.StatusCode)

		}
	},
}

func init() {
	addUserCmd.Flags().StringP("group", "g", "", "Name of the group (required)")
	addUserCmd.Flags().StringP("email", "e", "", "Email of the user to add (required)")
	addUserCmd.MarkFlagRequired("group")
	addUserCmd.MarkFlagRequired("email")
	GroupCmd.AddCommand(addUserCmd)
}
