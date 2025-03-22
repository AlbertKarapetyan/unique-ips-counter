package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	PrintMemUsage("Before start")

	file, err := os.Open("ip_addresses.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numWorkers := runtime.NumCPU()

	fmt.Printf("Number of logical CPUs: %d\n", numWorkers)

	lines := make(chan string, numWorkers)
	results := make(chan *ipSet, numWorkers)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(i, &wg, lines, results)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}

	close(lines)
	wg.Wait()
	close(results)

	fmt.Printf("Number of unique addresses: %d\n", NewIpSet().Count())
	fmt.Printf("Duration: %v\n", time.Since(startTime))
	PrintMemUsage("After end")
}
