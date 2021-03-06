package cmd

import (
	"fmt"
	"os"

	"github.com/justmiles/lambda-pagerduty-kasa/SmartPlug"
	"github.com/spf13/cobra"
)

var (
	username = os.Getenv("KASA_USERNAME")
	password = os.Getenv("KASA_PASSWORD")
)

var rootCmd = &cobra.Command{
	Use:     "kasa",
	Short:   "Interact with the Kasa's smart devices",
	Long:    `An unoffical tool to communicate with the remote Kasa API and control your smart devices`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute validates input the Cobra CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "your Kasa username (defaults to KASA_USERNAME environment variable)")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "your Kasa password (defaults to KASA_PASSWORD environment variable)")
	cobra.OnInitialize(loginToKasa)
}

func loginToKasa() {

	if username == "" {
		if os.Getenv("KASA_USERNAME") == "" {
			fmt.Println(`Please supply a username or export the "KASA_USERNAME" environment variable`)
			os.Exit(1)
		}
		username = os.Getenv("KASA_USERNAME")
	}

	if password == "" {
		if os.Getenv("KASA_PASSWORD") == "" {
			fmt.Println(`Please supply a password or export the "KASA_PASSWORD" environment variable`)
			os.Exit(1)
		}
		password = os.Getenv("KASA_PASSWORD")
	}

	smartplug.Login(username, password)
}
