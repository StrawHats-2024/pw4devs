/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Long:  `Register as new user. You can provide the email and password as arguments or, if not provided, the CLI will prompt you to enter them interactively.`,
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var email, password string
		if len(args) == 2 {
			email = args[0]
			password = args[1]
		} else {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Email").Value(&email),
					huh.NewInput().Title("Password").EchoMode(huh.EchoModePassword).Value(&password),
				),
			)
			form.Run()
		}
		fmt.Printf("register in with email: %s and password: %s\n", email, password)
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
