package group

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
)

// Delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a group",
	Long: `Remove a group from the system by its name.
You must provide the name of the group using the --name flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		groupName, _ := cmd.Flags().GetString("name")
		if groupName == "" {
			return fmt.Errorf("The --name flag is required")
		}
		type reqBody struct {
			GroupName string `json:"group_name"`
		}
		res, err := utils.MakeRequest[any]("/v1/groups", http.MethodDelete, reqBody{GroupName: groupName}, utils.GetAuthtoken())
		if err != nil {
			return err
		}
		switch res.StatusCode {
		case http.StatusNoContent:
			return nil
		case http.StatusNotFound:
			return fmt.Errorf("No group found with name: %s", groupName)
		default:
			return fmt.Errorf("Request failed with status code: %d", res.StatusCode)

		}
	},
}

func init() {
	deleteCmd.Flags().StringP("name", "n", "", "Name of the group to delete (required)")
	deleteCmd.MarkFlagRequired("name")
	GroupCmd.AddCommand(deleteCmd)
}
