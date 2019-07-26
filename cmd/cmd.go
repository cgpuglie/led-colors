package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var rootCmd = &cobra.Command{
	Use:   "led-colors",
	Short: "led-colors controls colors and effects on RGB leds",
}

// TODO: make these dynamic, add flags
var firmataAdaptor = firmata.NewAdaptor("/dev/ttyACM0")
var rgb = gpio.NewRgbLedDriver(firmataAdaptor, "11", "10", "9")

func init() {
	rootCmd.AddCommand(colorCmd)
	rootCmd.AddCommand(effectCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}