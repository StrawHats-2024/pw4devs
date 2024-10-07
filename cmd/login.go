/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login [email] [password]",
	Short: "Log in to your account",
	Long:  `Authenticates the user with their credentials. You can provide the email and password as arguments or, if not provided, the CLI will prompt you to enter them interactively.`,
	Args:  cobra.MaximumNArgs(2), // Accepts up to 2 arguments, or none for interactive input
	RunE: func(cmd *cobra.Command, args []string) error {
		var email, password string
		if len(args) == 2 {
			email = args[0]
			password = args[1]
		} else {
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
		switch response.StatusCode {
		case http.StatusOK:
			fmt.Printf("token %s", response.ResBody.Token)
			fmt.Printf("Login successfull!")
		case http.StatusUnauthorized:
			fmt.Printf("Unauthorized: %s", response.ResBody.Error)

		default:
			return fmt.Errorf("Unknow error encontered")
		}
		return nil
	},
}

func init() {
	authCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
