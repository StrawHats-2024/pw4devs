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

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from your account",
	Long:  "Ends the current authenticated session, logging the user out of their password manager account.",
	Run: func(cmd *cobra.Command, args []string) {

		if viper.GetString(utils.AuthTokenKey) != "" {
			utils.SetAuthToken("")
			fmt.Printf("Logout successful!")
			return
		} else {
			fmt.Println("User not authenticated")
		}
	},
}

func init() {
	authCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
