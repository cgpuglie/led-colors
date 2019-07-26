package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func pToV(percent float32) byte {
	if percent > 1 {
		return byte(254)
	}
	if percent < 0 {
		return byte(0)
	}
	return byte(int(percent * 254))
}

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	rgb := gpio.NewRgbLedDriver(firmataAdaptor, "11", "10", "9")

	// declare some consts
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

	robot.Start()
}
