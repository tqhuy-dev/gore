package utilities

import (
	"runtime"
)

type MemStats struct {
	HeapAlloc  uint64
	HeapSys    uint64
	StackInuse uint64
	NumGC      uint32
}

func CalMem() MemStats {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return MemStats{
		HeapAlloc:  memStats.HeapAlloc / 1024,
		HeapSys:    memStats.HeapSys / 1024,
		StackInuse: memStats.StackInuse / 1024,
		NumGC:      memStats.NumGC,
	}
}
