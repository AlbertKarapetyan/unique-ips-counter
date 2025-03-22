package main

import (
	"fmt"
	"runtime"
)

func PrintMemUsage(title string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s — Memory: %.2f MB\n", title, float64(m.Alloc)/1024/1024)
}
