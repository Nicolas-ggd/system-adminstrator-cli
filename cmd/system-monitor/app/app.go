package app

import (
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/monitor"
	"github.com/fatih/color"
	"os"
	"runtime"
	"time"
)

var (
	processing = color.New(color.Bold, color.FgGreen)
	invalid    = color.New(color.Bold, color.FgRed)
)

func Run() {
	arg := os.Args[1:]

	if arg[0] == "-info" {
		monitor.CpuLogger()
		os.Exit(0)
	}

	processing.Println("➜ system-monitor starting...")
	if arg[0] != "run" {
		help()
		os.Exit(0)
	}

	systemOs := detectOS()

	switch systemOs {
	case "linux":
		startStats, err := monitor.ReadCPUTasks()
		if err != nil {
			invalid.Printf("Error reading CPU stats, failed to %s\n", err.Error())
			invalid.Printf("Invalid OS system, your current OS is: %s\n", systemOs)
			os.Exit(0)
		}

		for {
			time.Sleep(1 * time.Second)
			endStats, err := monitor.ReadCPUTasks()
			if err != nil {
				invalid.Printf("Error reading CPU stats: %s\n", err.Error())
			}

			cpuUsage := monitor.CalculateCPUUsage(startStats, endStats)

			processing.Printf("➜ CPU Usage: %.2f%%\n", cpuUsage)
		}
	default:
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
	osSystem := runtime.GOOS

	switch osSystem {
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
		processing.Printf("%s.\n", osSystem)
	}

	return osSystem
}
