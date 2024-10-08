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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Long: `Create a new user account with an email and password. You can provide the email and password
as command flags, or if no flags are provided, an interactive form will be presented for input.`,
	Args: cobra.NoArgs, // No positional arguments, we use flags instead
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if the user is already authenticated
		if viper.GetString(utils.AuthTokenKey) != "" {
			fmt.Println("Already authenticated.")
			return nil
		}

		// Fetch flag values
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		// Check if both email and password are provided
		if email == "" || password == "" {
			// Show interactive form for input if flags are not provided
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

		// Send register request
		type RegisterResponse struct {
			Message string `json:"message"`
		}
		type RegisterBody struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		response, err := utils.MakeRequest[RegisterResponse]("/v1/auth/register",
			http.MethodPost, RegisterBody{Email: email, Password: password}, "")
		if err != nil {
			return err
		}

		// Handle response based on status code
		switch response.StatusCode {
		case http.StatusCreated:
			fmt.Println("Registration successful!")
		case http.StatusUnprocessableEntity:
			fmt.Println("Invalid email or password.")
		case http.StatusConflict:
			fmt.Println("Email already taken.")
		default:
			return fmt.Errorf("Unknown error encountered")
		}

		return nil
	},
}

// Add flags to the command
func init() {
	AuthCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("email", "e", "", "New user's email address (required)")
	registerCmd.Flags().StringP("password", "p", "", "New user's password (required)")

}
