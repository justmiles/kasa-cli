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
	outputType string
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&outputType, "output", "o", "text", `output type "text" or "json"`)
}

// process the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list devices in kasa",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the devices from Kasa
		devices := smartplug.GetDeviceList()

		// output a JSON response if requested
		if outputType == "json" {
			devicesJSON, _ := json.MarshalIndent(devices, "", "  ")
			fmt.Println(string(devicesJSON))

			// output a pretty table
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Alias", "Status", "Name", "Model", "MAC Address", "Device ID"})

			for _, device := range devices {
				var deviceStatus string
				if device.Status == 1 {
					deviceStatus = "Online"
				} else {
					deviceStatus = "Offline"
				}
				table.Append([]string{device.Alias, deviceStatus, device.DeviceName, device.DeviceModel, device.DeviceMac, device.DeviceID})
			}

			table.Render()
		}
	},
}
