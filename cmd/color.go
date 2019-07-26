package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"time"
	"gobot.io/x/gobot"
)

const (
	full = byte(254)
	half = byte(int(full / 2))
	off = byte(0)
)
var colorCmd = &cobra.Command{
	Use: "color",
	Short: "Set an RGB color",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Must supply a color argument")
			os.Exit(1)
		}
	
		work := func() {
			gobot.Every(1*time.Second, func() {
				switch args[0] {
				case "red":
					rgb.SetRGB(full, off, off)
				case "green":
					rgb.SetRGB(off, full, off)
				case "blue":
					rgb.SetRGB(off, off, full)
				case "purple":
					rgb.SetRGB(half, off, half)
				}
			})
		}
	
		robot := gobot.NewRobot("color",
			[]gobot.Connection{firmataAdaptor},
			[]gobot.Device{rgb},
			work,
		)
	
		robot.Start()
	},
}