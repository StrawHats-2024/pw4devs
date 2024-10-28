package dev

import (
	"github.com/spf13/cobra"
	"strawhats.pm4dev/internals/utils"
)

// SeedCmd represents the auth command
var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the db",
	Long:  "Only for development",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Command logic here
		return utils.SeedDb()
	},
}

func init() {
	// rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
