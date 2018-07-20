package cmd

import (
	"github.com/justmiles/lambda-pagerduty-kasa/SmartPlug"
	"github.com/spf13/cobra"
)

var (
	deviceAliasOff string
)

func init() {
	rootCmd.AddCommand(offCmd)
	offCmd.Flags().StringVarP(&deviceAliasOff, "device", "d", "text", `alias of device to turn off`)
	offCmd.MarkFlagRequired("device")

}

// process the toggle command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "turn off a device",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the devices from Kasa
		device := smartplug.GetDeviceByAlias(deviceAliasOff)
		device.Off()

	},
}
