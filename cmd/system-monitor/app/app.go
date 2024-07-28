package app

import (
	"github.com/fatih/color"
	"os"
	"runtime"
)

var (
	processing = color.New(color.Bold, color.FgGreen)
)

func Run() {
	processing.Println("➜ system-monitor starting...")

	_ = detectOS()

	arg := os.Args[1:]
	if arg[0] != "run" {
		help()
		os.Exit(0)
	}

}

func help() {
	c := color.New(color.FgRed)
	text := "Please provide a command to access system administrators."
	help := `
Usage
	$ system-monitor <run command>

Examples
	$ system-monitor run 
	`
	_, err := c.Printf("%s %s\n", text, help)
	if err != nil {
		return
	}

}

// detectOS is used to detect operating system and the target architecture of a running program.
func detectOS() string {
	os := runtime.GOOS

	switch os {
	case "windows":
		processing.Println("➜ OS: Windows System")
		processing.Printf("➜ Architecture: %s\n", runtime.GOARCH)
	case "darwin":
		processing.Println("➜ OS: Darwin")
		processing.Printf("➜ Architecture: %s\n", runtime.GOARCH)
	case "linux":
		processing.Println("➜ OS: Linux")
		processing.Printf("➜ Architecture: %s\n", runtime.GOARCH)
	default:
		processing.Printf("%s.\n", os)
	}

	return os
}
