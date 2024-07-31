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
			PID:         parse.ToInt(strings.TrimSpace(line[0:10])),
			Command:     strings.TrimSpace(line[11:61]),
			FullCommand: line[74:],
			CPU:         parse.ToFloat64(strings.TrimSpace(line[63:68])),
			Mem:         parse.ToFloat64(strings.TrimSpace(line[69:74])),
		}

		proc = append(proc, p)
	}

	return proc, nil
}
