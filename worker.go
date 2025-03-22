package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"sync"
)

func Worker(id int, wg *sync.WaitGroup, lines <-chan string, result chan<- *ipSet) {
	defer wg.Done()
	fmt.Printf("Goroutine id = %d\n", id)
	bs := NewIpSet()
	for line := range lines {
		ip := net.ParseIP(strings.TrimSpace(line))
		if ip == nil {
			continue
		}

		if IPv4 := ip.To4(); IPv4 != nil {
			num := binary.BigEndian.Uint32(IPv4)
			bs.Set(num)
		}
	}
	result <- bs
}
