package main

import (
	"github.com/cgpuglie/led-colors/cmd"
)

// currently assigns color code to RGB LEDs
// device, pins, and other config is static
// perhaps, there should be a 'server' process that is always running
// CLI can talk to the server through RPC, telling it what to do
// That way, we can keep the connection open and it will respond more quickly
func main() {
	cmd.Execute()
}
