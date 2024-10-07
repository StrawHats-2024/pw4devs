/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strawhats.pm4dev/internals/utils"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Check the current authentication status",
	Long:  "Displays the current authentication status of the user in the password manager. If authenticated, shows user details.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Command logic here
		if viper.GetString(utils.AuthTokenKey) != "" {
			fmt.Printf("Logged in as %s\n", viper.GetString(utils.UserEmailKey))
		} else {
			fmt.Printf(`User not authenticated.
Please perform one of the following actions:
  - auth login
  - auth register
`)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
