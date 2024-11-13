package group

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
)

// Get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve group information",
	Long: `Display details of a group by its name.
You must provide the name of the group using the --name flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		getUsers, _ := cmd.Flags().GetBool("users")
		type reqBody struct {
			GroupName string `json:"group_name"`
		}
		type resBody struct {
			Data    utils.GroupRecordWithUsers `json:"data"`
			Message string                     `json:"message"`
		}
		res, err := utils.MakeRequest[resBody]("/v1/groups", http.MethodGet, reqBody{GroupName: name}, utils.GetAuthtoken())
		if err != nil {
			return err
		}
		switch res.StatusCode {
		case http.StatusOK:
			if getUsers {
				prettyPrintGroupUsers(res.ResBody.Data)
			} else {
				utils.PrintSecrets(res.ResBody.Data.Secrets)
			}
			return nil
		default:
			return fmt.Errorf("Request failed with status code: %d", res.StatusCode)

		}
	},
}

func init() {
	getCmd.Flags().StringP("name", "n", "", "Name of the group to retrieve (required)")
	getCmd.Flags().BoolP("users", "u", false, "Retrieve users of the group")
	getCmd.MarkFlagRequired("name")
	GroupCmd.AddCommand(getCmd)
}

func prettyPrintGroupUsers(group utils.GroupRecordWithUsers) {
	totalUsers := len(group.Users)
	fmt.Printf("Total Users: %d\n", totalUsers)
	fmt.Println("User Emails:")

	for _, user := range group.Users {
		fmt.Println("\t" + user.Email)
	}
}
