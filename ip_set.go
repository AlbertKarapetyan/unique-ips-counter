package main

import (
	"sync"
)

var (
	ips  map[uint32]byte = make(map[uint32]byte)
	lock sync.Mutex
)

type ipSet struct {
}

func NewIpSet() *ipSet {
	return &ipSet{}
}

func (i *ipSet) Set(ip uint32) {
	lock.Lock()
	ips[ip] = 1
	lock.Unlock()
}

func (i *ipSet) Count() int {
	return len(ips)
}
