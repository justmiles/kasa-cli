package cmd

import (
	"github.com/justmiles/lambda-pagerduty-kasa/SmartPlug"
	"github.com/spf13/cobra"
)

var (
	deviceAliasOn string
)

func init() {
	rootCmd.AddCommand(onCmd)
	onCmd.Flags().StringVarP(&deviceAliasOn, "device", "d", "", `alias of device to turn on`)
	onCmd.MarkFlagRequired("device")

}

// process the toggle command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "turn on a device",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the devices from Kasa
		device := smartplug.GetDeviceByAlias(deviceAliasOn)
		device.On()

	},
}
