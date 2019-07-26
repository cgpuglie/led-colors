package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gobot.io/x/gobot"
)

const (
	full = byte(254)
	half = byte(int(full / 2))
	off  = byte(0)
)

// TODO: add argument info
var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "Set an RGB color",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Must supply a color argument")
			os.Exit(1)
		}

		work := func() {
			gobot.Every(1*time.Second, func() {
				// TODO: add more colors
				switch args[0] {
				case "red":
					rgb.SetRGB(full, off, off)
				case "green":
					rgb.SetRGB(off, full, off)
				case "blue":
					rgb.SetRGB(off, off, full)
				case "purple":
					rgb.SetRGB(full, off, full)
				}
			})
		}

		robot := gobot.NewRobot("color",
			[]gobot.Connection{firmataAdaptor},
			[]gobot.Device{rgb},
			work,
		)

		// TODO: does this half to hang? Can we upload the process instead and let it do it's thing?
		robot.Start()
	},
}
