package util

import (
	"fmt"
	"log"
	"net"
)

// GetLocalIPs Attempt to get local IPv4 address, currently only used for quick access web url in terminal
func GetLocalIPs() []net.IP {
	// Get Network Interface Addresses
	var ips []net.IP
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("[ERROR] Error getting local IP address: ", err.Error())
		log.Fatal(err)
	}

	// Discard loopback addresses
	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	return ips
}
