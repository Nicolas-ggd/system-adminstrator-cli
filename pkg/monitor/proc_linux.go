package monitor

import (
	"fmt"
	"github.com/Nicolas-ggd/system-adminstrator-cli/pkg/parse"
	"log"
	"os/exec"
	"strings"
)

type ProcMonitor struct {
	PID         int     `json:"pid"`
	Command     string  `json:"command"`
	FullCommand string  `json:"full_command"`
	CPU         float64 `json:"percent-cpu"`
	Mem         float64 `json:"percent-mem"`
}

func GetProc() ([]ProcMonitor, error) {
	var proc []ProcMonitor
	val, err := exec.Command("ps", "-axo", "pid:10,comm:50,pcpu:5,pmem:5,args").Output()
	if err != nil {
		log.Fatalln("fail to get process info")
	}

	procSettings := strings.Split(strings.TrimSuffix(string(val), "\n"), "\n")[1:]
	for _, line := range procSettings {
		fmt.Printf("%+v\n", line)

		p := ProcMonitor{
			// Characters 0 to 10 (PID)
			PID: parse.ToInt(strings.TrimSpace(line[0:10])),
			// Characters 11 to 61 (Command)
			Command: strings.TrimSpace(line[11:61]),
			// Characters ti 74 and > is (full command)
			FullCommand: line[74:],
			// Characters 63 to 68 (CPU usage)
			CPU: parse.ToFloat64(strings.TrimSpace(line[63:68])),
			// Characters 69 to 74 (Memory usage)
			Mem: parse.ToFloat64(strings.TrimSpace(line[69:74])),
		}

		proc = append(proc, p)
	}

	return proc, nil
}
