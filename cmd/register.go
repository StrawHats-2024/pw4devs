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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Long:  `Register as new user. You can provide the email and password as arguments or, if not provided, the CLI will prompt you to enter them interactively.`,
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		var email, password string
		if len(args) == 2 {
			email = args[0]
			password = args[1]
		} else {
			// TODO: add repeat password
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
		type RegisterResponse struct {
			Message string `json:"message"`
		}
		type LoginBody struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		response, err := utils.MakeRequest[RegisterResponse]("/v1/auth/register",
			http.MethodPost, LoginBody{Email: email, Password: password}, "")
		if err != nil {
			return err
		}
		switch response.StatusCode {
		case http.StatusCreated:
			fmt.Printf("Register successfull!")
		case http.StatusUnprocessableEntity:
			fmt.Printf("Invalid email or password")
		case http.StatusConflict:
			fmt.Printf("Email already taken")
		default:
			return fmt.Errorf("Unknow error encontered")
		}
		return nil

	},
}

func init() {
	authCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
