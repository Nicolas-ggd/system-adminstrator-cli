// Copyright (c) 2024 Nicolas-ggd, released under Apache-2.0 License. See LICENSE file.

package app

import (
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/cli"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/monitor"
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	processing = color.New(color.FgGreen)
	invalid    = color.New(color.Bold, color.FgRed)
)

func Run() {
	systemOs := detectOS()

	arg := os.Args[1:]
	processing.Println("➜ system-monitor starting...")

	if len(arg) == 0 {
		invalid.Println("No arguments provided.")
		os.Exit(0)
	}

	switch arg[0] {
	case "run":
		switch systemOs {
		case "linux":
			// count cpu core with logical cores
			cpuCount := monitor.CountCPUCore()

			// read cpu tasks
			startStats, err := monitor.ReadCPUTasks(cpuCount)
			if err != nil {
				invalid.Printf("Error reading CPU stats, failed to %s\n", err.Error())
				invalid.Printf("Invalid OS system, your current OS is: %s\n", systemOs)
				os.Exit(0)
			}

			// read network stats
			startNetStat, _ := monitor.ReadNetDev()

			for {
				// wait 1 second for delay
				time.Sleep(1 * time.Second)

				// read cpu tasks again after 1-second delay
				endStats, err := monitor.ReadCPUTasks(cpuCount)
				if err != nil {
					invalid.Printf("Error reading CPU stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}

				// read networks tasks again after 1-second delay
				endNetStat, err := monitor.ReadNetDev()
				if err != nil {
					invalid.Printf("Error reading net stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}
				// clear screen
				cli.ClearScreen()

				// receive calculated network usage
				netStat, err := monitor.ReadNetUsage(startNetStat, endNetStat)
				if err != nil {
					invalid.Printf("Error reading Network stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}

				fmt.Printf("%+v\n", netStat)

				// read memory usage
				memResp, err := monitor.ReadMemUsage()
				if err != nil {
					invalid.Printf("Error reading Memory and Swap stats: %s\n", err.Error())
					continue // skip this iteration and retry in the next loop
				}

				// calculate cpu usage
				cpuUsage, err := monitor.CalculateCPUUsage(startStats, endStats)
				if err != nil {
					log.Fatalln(err)
				}

				processing.Printf("system-monitoring - %v", time.Now().Format("15:04:05"))
				// draw table
				table := cli.DrawTable(cpuUsage, memResp)

				// render table
				table.Render()

				// update startStats to endStats for the next interval calc
				startStats = endStats
			}
		default:
			os.Exit(0)
		}
	case "info":
		monitor.CpuLogger()
		os.Exit(0)
	case "help":
		help()
	case "proc":
		switch systemOs {
		case "linux":
			// wait 1 second for delay
			time.Sleep(1 * time.Second)

			// clear the terminal screen
			cli.ClearScreen()

			processes, err := monitor.GetProc()
			if err != nil {
				log.Fatalf("Error fetching process info: %v", err)
			}

			fmt.Printf("%-5s %-25s %-60s %-8s %-8s\n", "PID", "COMMAND", "ARGS", "CPU%", "MEM%")
			fmt.Println(strings.Repeat("-", 110))

			// iterate through processes and print each one
			for _, proc := range processes {
				fmt.Printf("%-5d %-25s %-60s %-8.1f %-8.1f\n", proc.PID, proc.Command, proc.FullCommand, proc.CPU, proc.Mem)
			}
		}
	default:
		invalid.Println("➜ Please provide a command to access system administrators or run command `help`")
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
	$ system-monitor proc 
	$ system-monitor info 
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
	case "linux":
		processing.Printf("➜ OS: %s\n", osSystem)
		processing.Printf("➜ Architecture: %s\n", osSystem)
	default:
		processing.Printf("%s.\n", osSystem)
	}

	return osSystem
}
