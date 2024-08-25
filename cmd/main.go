package main

import (
	"fmt"
	"log"

	"github.com/kalom60/Hardware-Monitor/hardware"
)

func main() {
    systemSection, err := hardware.GetSystemInfo()
    if err != nil {
        log.Fatal("System Section Error", err.Error())
        return
    }

    fmt.Println("System Section: ", systemSection)

    diskSection, err := hardware.GetDiskInfo()
    if err != nil {
        log.Fatal("Disk Section Error", err.Error())
        return
    }

    fmt.Println("Disk Section: ", diskSection)

    cpuSection, err := hardware.GetCPUInfo()
    if err != nil {
        log.Fatal("CPU Section Error", err.Error())
        return
    }

    fmt.Println("CPU Section: ", cpuSection)
}
