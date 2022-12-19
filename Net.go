package main

import "net"

func getIfaceNames() (result []string, err error) {
	ifaces, err := net.Interfaces()
	result = make([]string, len(ifaces))
	for i := 0; i < len(ifaces); i++ {
		result[i] = ifaces[i].Name
	}
	return
}
