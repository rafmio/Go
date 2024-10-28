// GC - Garbage Collection
package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.Sys:", mem.Sys)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----------------")
	fmt.Println("mem.HeapSys:", mem.HeapSys)
	fmt.Println("mem.HeapIdle:", mem.HeapIdle)
	fmt.Println("mem.HeapInuse:", mem.HeapInuse)
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50_000_000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}

	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100_000_000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(100 * time.Millisecond)
	}
	printStats(mem)
}
