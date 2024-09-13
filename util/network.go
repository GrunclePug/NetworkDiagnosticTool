package util

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type NetInfo struct {
	Interfaces []net.IP
	PublicIP   net.IP
}

func GetNetInfo() NetInfo {
	return NetInfo{
		Interfaces: GetLocalIPs(),
		PublicIP:   GetPublicIP(),
	}
}

// GetLocalIPs Attempt to get local IPv4 address, currently only used for quick access web url in terminal
func GetLocalIPs() []net.IP {
	// Get Network Interface Addresses
	var ips []net.IP
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("[ERROR] Error getting local IP address: ", err.Error())
		log.Fatal(err)
	}

	// Discard loopback & APIPA
	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && !strings.HasPrefix(ipnet.IP.String(), "169") {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	return ips
}

func GetPublicIP() net.IP {
	resp, err := http.Get("https://ident.me/")
	if err != nil {
		fmt.Println("[ERROR] Error getting public IP address: ", err.Error())
		log.Fatal(err)
	}

	content, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("[ERROR] Error getting public IP address: ", err.Error())
		log.Fatal(err)
	}
	return net.ParseIP(string(content))
}
