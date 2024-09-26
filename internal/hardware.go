package hardware

import (
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

const mbDiv uint64 = 1024 * 1024
const gbDiv uint64 = mbDiv * 1024

func GetSystemInfo() (string, error) {
	runtimeOS := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	hostStat, err := host.Info()
	if err != nil {
		return "", err
	}

	html := "\n"
	html = html + "\tOperating System:" + runtimeOS + "\n"
	html = html + "\tPlatform:" + hostStat.Platform + "\n"
	html = html + "\tHostname:" + hostStat.Hostname + "\n"
	html = html + "\tNumber of processes running:" + strconv.FormatUint(hostStat.Procs, 10) + "\n"
	html = html + "\tTotal memory:" + strconv.FormatUint(vmStat.Total/mbDiv, 10) + " MB" + "\n"
	html = html + "\tFree memory:" + strconv.FormatUint(vmStat.Free/mbDiv, 10) + " MB" + "\n"
	html = html + "\tPercentage used memory:" + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%" + "\n"

	return html, nil
}

func GetDiskInfo() (string, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return "", err
	}

	html := "\n"
	html = html + "\tTotal disk space:" + strconv.FormatUint(diskStat.Total/gbDiv, 10) + " GB" + "\n"
	html = html + "\tUsed disk space:" + strconv.FormatUint(diskStat.Used/gbDiv, 10) + " GB" + "\n"
	html = html + "\tFree disk space:" + strconv.FormatUint(diskStat.Free/gbDiv, 10) + " GB" + "\n"
	html = html + "\tPercentage disk space usage:" + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%" + "\n"
	return html, nil
}

func GetCpuInfo() (string, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return "", err
	}
	percentage, err := cpu.Percent(0, true)
	if err != nil {
		return "", err
	}

	html := "\n"

	if len(cpuStat) != 0 {
		html = html + "\tModel Name:" + cpuStat[0].ModelName + "\n"
		html = html + "\tFamily:" + cpuStat[0].Family + "\n"
		html = html + "\tSpeed:" + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz" + "\n"
	}

	firstCpus := percentage[:len(percentage)/2]
	secondCpus := percentage[len(percentage)/2:]

	html = html + "Cores: \n"
	for idx, cpupercent := range firstCpus {
		html = html + "\tCPU [" + strconv.Itoa(idx) + "]: " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%" + "\n"
	}
	html = html + ""
	for idx, cpupercent := range secondCpus {
		html = html + "\tCPU [" + strconv.Itoa(idx+8) + "]: " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%" + "\n"
	}
	return html, nil
}
