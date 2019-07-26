package main

import (

	"github.com/cgpuglie/led-colors/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "led-colors",
	Short: "led-colors controls colors and effects on RGB leds",
}

func main() {
	cmd.Execute()
}
