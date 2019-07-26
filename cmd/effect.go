package cmd

import (
	"github.com/spf13/cobra"
	"time"
	"gobot.io/x/gobot"
)

var effectCmd = &cobra.Command{
	Use: "effect",
	Short: "Set an RGB effect",
}

var cycleCmd = &cobra.Command{
	Use: "cycle",
	Short: "Cycle between all LED colors",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: clean this up, copied from a standalone script
		// declare some consts
		// TODO: provide flags for these
		const (
			min = 0
			max = 100
			ms  = 20
		)
	
		// declare some vars
		var (
			r = true
			b = false
			g = false
			c = 0
			p = max
		)
	
		work := func() {
			gobot.Every(ms*time.Millisecond, func() {
				// increment/decrement
				c++
				p--
				if r {
					// handle red
					rgb.SetRGB(byte(c), byte(min), byte(p))
					if c == max {
						c = min
						p = max
						r = false
						g = true
					}
				} else if g {
					// handle green
					rgb.SetRGB(byte(p), byte(c), byte(min))
					if c == max {
						c = min
						p = max
						g = false
						b = true
					}
				} else if b {
					// handle bue
					rgb.SetRGB(byte(min), byte(p), byte(c))
					if c == max {
						c = min
						p = max
						b = false
						r = true
					}
				}
			})
		}
	
		robot := gobot.NewRobot("bot",
			[]gobot.Connection{firmataAdaptor},
			[]gobot.Device{rgb},
			work,
		)
		
		// TODO: does this half to hang? Can we upload the process instead and let it do it's thing?
		robot.Start()
	},
}