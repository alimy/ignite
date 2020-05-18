package main

import (
	"github.com/alimy/ignite/cmd"
)

func main() {
	cmd.Setup(
		"ignite",          // command name
		"vm help toolkit", // command short describe
		"vm help toolkit", // command long describe
	)
	// execute start application
	cmd.Execute()
}
