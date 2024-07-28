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
)

func Run() {
	arg := os.Args[1:]

	if arg[0] == "-info" {
		monitor.CpuLogger()
	}

	processing.Println("➜ system-monitor starting...")
	if arg[0] != "run" && arg[0] != "-info" {
		help()
		os.Exit(0)
	}

	systemOs := detectOS()

	switch systemOs {
	case "linux":
		for {
			idle0, total0 := monitor.GetLinuxCPU()
			time.Sleep(3 * time.Second)
			idle1, total1 := monitor.GetLinuxCPU()

			idleTicks := float64(idle1 - idle0)
			totalTicks := float64(total1 - total0)
			cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

			processing.Printf("➜  CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
		}
	case "darwin":
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
