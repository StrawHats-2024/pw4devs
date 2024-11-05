package sharing

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
)

// ToGroup command
var toGroupCmd = &cobra.Command{
	Use:   "togroup",
	Short: "Share a secret with a group",
	Long: `Grant access to a secret for all members of a specified group.
You must provide the ID of the secret and the name of the group using the --secret-id and --group flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		secretID, _ := cmd.Flags().GetString("secret-id")
		group, _ := cmd.Flags().GetString("group")
		if secretID == "" || group == "" {
			return fmt.Errorf("Both --secret-id and --group flags are required")
		}

		id, err := strconv.Atoi(secretID)
		if err != nil {
			return err
		}
		type reqBody struct {
			SecretID   int    `json:"secret_id"`
			GroupName  string `json:"group_name"`
			Permission string `json:"permission"`
		}
		res, err := utils.MakeRequest[any]("/v1/secrets/share/group", http.MethodPost, reqBody{
			SecretID:   id,
			GroupName:  group,
			Permission: "read-only",
		}, utils.GetAuthtoken())
		if err != nil {
			return err
		}
		switch res.StatusCode {
		case http.StatusCreated:
			fmt.Println("res: ", res.ResBody)
			return nil
		default:
			fmt.Println("res: ", res.ResBody)
			return fmt.Errorf("Request failed with response code %d", res.StatusCode)

		}
	},
}

func init() {
	toGroupCmd.Flags().StringP("secret-id", "s", "", "ID of the secret to share (required)")
	toGroupCmd.Flags().StringP("group", "g", "", "Name of the group to share with (required)")
	toGroupCmd.MarkFlagRequired("secret-id")
	toGroupCmd.MarkFlagRequired("group")
	ShareCmd.AddCommand(toGroupCmd)
}
