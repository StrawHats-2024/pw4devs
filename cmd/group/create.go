package group

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
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
		type reqBody struct {
			GroupName string `json:"group_name"`
		}
		res, err := utils.MakeRequest[any]("/v1/groups", http.MethodPost, reqBody{GroupName: name}, utils.GetAuthtoken())
		if err != nil {
			return err
		}
		switch res.StatusCode {
		case http.StatusCreated:
			// fmt.Println("res: ", res.ResBody)
			return nil
		case http.StatusConflict:
			return fmt.Errorf("Group name already taken")
		default:
			return fmt.Errorf("Request failed with status code: %d", res.StatusCode)

		}
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "Name of the new group (required)")
	createCmd.MarkFlagRequired("name")
	GroupCmd.AddCommand(createCmd)
}
