package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	username string
	password string
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
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "your Kasa username")
	rootCmd.PersistentFlags().StringVarP(&username, "password", "p", "", "your Kasa password")
}
