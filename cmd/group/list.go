package group

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
)

// List command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all groups",
	Long: `Display a list of all groups accessible to the current user.
You can specify the maximum number of groups to display and the number of groups to skip using the --limit and --offset flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		type resBody struct {
			Data    []utils.GroupRecord `json:"data"`
			Message string            `json:"message"`
		}

    //TODO: Fix this 
		res, err := utils.MakeRequest[resBody]("/v1/groups/user", http.MethodGet, nil, utils.GetAuthtoken())
		if err != nil {
			return err
		}
		switch res.StatusCode {
		case http.StatusOK:
      for _, group := range res.ResBody.Data {
        fmt.Println(group.Name)
      }
			return nil
		default:
			return fmt.Errorf("Request failed with status code: %d", res.StatusCode)

		}
	},
}

func init() {
	GroupCmd.AddCommand(listCmd)
}
