/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package auth

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strawhats.pm4dev/internals/utils"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the system",
	Long: `Authenticate a user with their email and password. You can provide the email and password
as command flags, or if no flags are provided, an interactive form will be presented for input.`,
	Args: cobra.NoArgs, // No positional arguments, we use flags instead
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString(utils.AuthTokenKey) != "" {
			fmt.Println("Already authenticated.")
			return nil
		}

		// Fetch flag values
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		// Check if both email and password are provided
		if email == "" || password == "" {
			// If not, show interactive form for input
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Email").Value(&email).
						Validate(utils.ValidateEmail),
					huh.NewInput().Title("Password").EchoMode(huh.EchoModePassword).Value(&password).
						Validate(utils.ValidatePassword),
				),
			)
			form.Run()
		}

		// Send login request
		type LoginResponse struct {
			Token   string `json:"token"`
			Error   string `json:"error"`
			Message string `json:"message"`
		}
		type LoginBody struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		response, err := utils.MakeRequest[LoginResponse]("/v1/auth/login",
			http.MethodPost, LoginBody{Email: email, Password: password}, "")
		if err != nil {
			return err
		}

		// Handle response based on status code
		switch response.StatusCode {
		case http.StatusOK:
			authToken := response.ResBody.Token
			if authToken == "" {
				return fmt.Errorf("No auth token received.")
			}
			// Save the auth token and user email
			if err := utils.SetAuthToken(authToken); err != nil {
				return err
			}
			if err := utils.SetUserEmail(email); err != nil {
				return err
			}
			fmt.Println("Login successful!")
		case http.StatusUnauthorized:
			fmt.Printf("Unauthorized: %s\n", response.ResBody.Error)
		default:
			return fmt.Errorf("Unknown error encountered")
		}

		return nil
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "User's email address (required)")
	loginCmd.Flags().StringP("password", "p", "", "User's password (required)")

	// Mark email and password as required flags
	// loginCmd.MarkFlagRequired("email")
	// loginCmd.MarkFlagRequired("password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
