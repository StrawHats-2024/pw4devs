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
		if name == "" {
			return fmt.Errorf("The --name flag is required")
		}
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
      //TODO: impliment pretty print
			fmt.Printf("res: %+v", res.ResBody.Data)
			return nil
		default:
			return fmt.Errorf("Request failed with status code: %d", res.StatusCode)

		}
	},
}

func init() {
	getCmd.Flags().StringP("name", "n", "", "Name of the group to retrieve (required)")
	getCmd.MarkFlagRequired("name")
	GroupCmd.AddCommand(getCmd)
}
