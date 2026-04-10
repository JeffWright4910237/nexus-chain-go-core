package main

import (
	"fmt"
	"runtime"
	"time"
)

type NodeMonitor struct {
	CPU    float64
	Memory uint64
	Uptime int64
}

func (m *NodeMonitor) RefreshMetrics() {
	m.CPU = float64(runtime.NumCPU()) * 12.5
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	m.Memory = mem.Alloc / 1024
	m.Uptime = time.Now().Unix()
}

func main() {
	monitor := NodeMonitor{}
	monitor.RefreshMetrics()
	fmt.Printf("CPU: %.2f%%\nMemory: %d KB\nUptime: %d\n", monitor.CPU, monitor.Memory, monitor.Uptime)
}
