package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/justmiles/lambda-pagerduty-kasa/SmartPlug"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	deviceAliasStatus string
)

func init() {
	rootCmd.AddCommand(statusCmd)
	// statusCmd.MarkFlagRequired("device")
}

// process the toggle command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "get device status",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the devices from Kasa
		device := smartplug.GetDeviceByAlias(deviceAliasStatus)
		info := device.GetSystemInfo()

		if outputType == "json" {
			infoJSON, _ := json.Marshal(info)
			fmt.Println(string(infoJSON))
		} else {

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Alias", "Relay State", "Name", "Model", "MAC Address", "Device ID"})

			var relayState string
			if info.RelayState == 1 {
				relayState = "On"
			} else {
				relayState = "Off"
			}
			table.Append([]string{info.Alias, relayState, info.DevName, info.Model, info.Mac, info.DeviceID})

			table.Render()

		}

	},
}
