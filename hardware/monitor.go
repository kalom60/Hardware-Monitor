package hardware

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

const (
	megabyte uint64 = 1024 * 1024
	gigabyte uint64 = megabyte * 1024
)

type System struct {
	OS                   string `json:"os"`
	Platform             string `json:"platform"`
	Hostname             string `json:"hostname"`
	Procs                string `json:"procs"`
	TotalMemory          string `json:"totalMemory"`
	FreeMemory           string `json:"freeMemory"`
	PercentageUsedMemory string `json:"percentageUsedMemory"`
}

func GetSystemInfo() (System, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return System{}, err
	}

	hostInfo, err := host.Info()
	if err != nil {
		return System{}, err
	}

	return System{
		OS:                   hostInfo.OS,
		Platform:             hostInfo.Platform,
		Hostname:             hostInfo.Hostname,
		Procs:                strconv.FormatUint(hostInfo.Procs, 10),
		TotalMemory:          strconv.FormatUint(vmStat.Total/megabyte, 10),
		FreeMemory:           strconv.FormatUint(vmStat.Free/megabyte, 10),
		PercentageUsedMemory: strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64),
	}, nil
}

type Disk struct {
	TotalDS           string `json:"totalDS"`
	UsedDS            string `json:"usedDS"`
	FreeDS            string `json:"freeDS"`
	PercentageDSUsage string `json:"percentageDSUsage"`
}

func GetDiskInfo() (Disk, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return Disk{}, err
	}

	return Disk{
		TotalDS:           strconv.FormatUint(diskStat.Total/gigabyte, 10),
		UsedDS:            strconv.FormatUint(diskStat.Used/gigabyte, 10),
		FreeDS:            strconv.FormatUint(diskStat.Free/gigabyte, 10),
		PercentageDSUsage: strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64),
	}, nil
}

type CPU struct {
	ModelName string `json:"modelName"`
	Family    string `json:"family"`
	Speed     string `json:"speed"`
	Cores     any    `json:"cors"`
}

func GetCPUInfo() (CPU, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return CPU{}, err
	}

    percentage, err := cpu.Percent(0, true)
    if err != nil {
        return CPU{}, err
    }

    fmt.Println("Percentage man", percentage)

	if len(cpuStat) != 0 {
		return CPU{
			ModelName: cpuStat[0].ModelName,
			Family:    cpuStat[0].Family,
			Speed:     strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64),
		}, nil
	}

	return CPU{}, nil
}
