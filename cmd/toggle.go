package cmd

import (
	"github.com/justmiles/lambda-pagerduty-kasa/SmartPlug"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(toggleCmd)
	toggleCmd.MarkFlagRequired("device")
}

// process the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "toggle device state",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the devices from Kasa
		device := smartplug.GetDeviceByAlias(deviceAlias)
		device.Toggle()

	},
}
