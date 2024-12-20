/*
Copyright © 2024 parikshith078

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"dv/cmd/auth"
	"dv/cmd/dev"
	"dv/cmd/group"
	"dv/cmd/secrets"
	"dv/cmd/sharing"
	"dv/internals/ui"
	"dv/internals/utils"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dv",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runnig")
		if auth := utils.GetAuthtoken(); auth == "" {
			os.Exit(1)
		} // check for auth status
		ui.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dv.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(secrets.SecretsCmd)
	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(group.GroupCmd)
	rootCmd.AddCommand(sharing.ShareCmd)
	rootCmd.AddCommand(dev.SeedCmd)
	rootCmd.AddCommand(completionCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".dv" (without extension).
		viper.AddConfigPath(home)
		// viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".pm4devs.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in or create one.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found, creating a new one...")
			err := viper.SafeWriteConfigAs(".pm4devs.yaml")
			if err != nil {
				cobra.CheckErr(err)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Error reading config file:", err)
		}
	} else {
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
