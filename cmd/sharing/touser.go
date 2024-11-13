package sharing

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"dv/internals/utils"
)

// ToUser command
var toUserCmd = &cobra.Command{
	Use:   "touser",
	Short: "Share a secret with a user",
	Long: `Grant access to a secret for a specific user by their email address.
You must provide the ID of the secret and the user's email address using the --secret-id and --email flags.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		secretID, _ := cmd.Flags().GetString("secret-id")
		email, _ := cmd.Flags().GetString("email")
		if secretID == "" || email == "" {
			return fmt.Errorf("Both --secret-id and --email flags are required")
		}
		id, err := strconv.Atoi(secretID)
		if err != nil {
			return err
		}
		type reqBody struct {
			SecretID   int    `json:"secret_id"`
			UserEmail  string `json:"user_email"`
			Permission string `json:"permission"`
		}
		// Placeholder for share to user logic
		res, err := utils.MakeRequest[any]("/v1/secrets/share/user", http.MethodPost, reqBody{
			SecretID:   id,
			UserEmail:  email,
			Permission: "read-only",
		}, utils.GetAuthtoken())

		switch res.StatusCode {
		case http.StatusCreated:
			return nil

		default:
			fmt.Println("res: ", res.ResBody)
			return fmt.Errorf("Request failed with response code %d", res.StatusCode)
		}
	},
}

func init() {
	toUserCmd.Flags().StringP("secret-id", "s", "", "ID of the secret to share (required)")
	toUserCmd.Flags().StringP("email", "e", "", "Email of the user to share with (required)")
	toUserCmd.MarkFlagRequired("secret-id")
	toUserCmd.MarkFlagRequired("email")
	ShareCmd.AddCommand(toUserCmd)
}
