package main

import (
	"fmt"
	hardware "go-sys-monitoring/internal"
	"time"
)

func main() {
	systemData, err := hardware.GetSystemInfo()
	if err != nil {
		fmt.Println(err)
	}
	diskData, err := hardware.GetDiskInfo()
	if err != nil {
		fmt.Println(err)
	}
	cpuData, err := hardware.GetCpuInfo()
	if err != nil {
		fmt.Println(err)
	}
	timeStamp := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println("Time: " + timeStamp)
	fmt.Println("System Data: " + systemData)
	fmt.Println("CPU Data: " + cpuData)
	fmt.Println("Disk Data: " + diskData)
}
