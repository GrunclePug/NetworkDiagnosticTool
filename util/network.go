package util

import (
	"fmt"
	"github.com/jackpal/gateway"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type NetInfo struct {
	IPs		   []net.IP
	Subnets	   []string
	Interfaces []string
	Macs 	   []string
	Gateway    net.IP
	PublicIP   net.IP
}

func GetNetInfo() NetInfo {
	return NetInfo{
		IPs: 	    LocalIPs(),
		Subnets:	subnets(),
		Interfaces: interfaces(),
		Macs:		macs(),
		Gateway:	router(),
		PublicIP:   publicIP(),
	}
}

// LocalIPs Attempt to get local IPv4 address, currently only used for quick access web url in terminal
func LocalIPs() []net.IP {
	var ips []net.IP

	// Get a list of network interface addresses
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("[ERROR] Error getting local IP address: ", err.Error())
		log.Fatal(err)
	}

	// Iterate through the addresses and discard loopback & APIPA while collecting the rest
	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && !strings.HasPrefix(ipnet.IP.String(), "169") {
				ips = append(ips, ipnet.IP)
			}
		}
	}

	return ips
}

// subnets Attempt to get network subnets
func subnets() []string {
	var subnets []string

	// Get a list of network interface addresses
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("[ERROR] Error getting local IP address for subnets: ", err.Error())
		log.Fatal(err)
	}

	// Discard loopback & APIPA
	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && !strings.HasPrefix(ipnet.IP.String(), "169") {
				// Calculate the network address by bitwise AND between IP and subnet mask
				networkIP := ipnet.IP.Mask(ipnet.Mask)

				// Get the CIDR size (number of leading 1's in the mask)
				ones, _ := ipnet.Mask.Size()

				// Append the network address and CIDR size
				subnet := fmt.Sprintf("%s/%d", networkIP.String(), ones)
				subnets = append(subnets, subnet)
			}
		}
	}

	return subnets
}

// interfaces Attempt to get all network interfaces
func interfaces() []string {
	var interfaceNames []string

	// Get a list of network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("[ERROR] Unable to get network interfaces:", err.Error())
		return interfaceNames
	}

	// Iterate through the interfaces and collect their names
	for _, iface := range interfaces {
		interfaceNames = append(interfaceNames, iface.Name)
	}

	return interfaceNames
}

// macs Attempt to get all mac addresses for device
func macs() []string {
	var macs []string

	// Get a list of network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("[ERROR] Unable to get network interfaces for MAC: ", err.Error())
		log.Fatal(err)
		return macs
	}

	// Iterate through the interface and collect their MAC addresses
	for _, iface := range interfaces {
		mac := iface.HardwareAddr.String()
		if mac != "" {
			macs = append(macs, mac)
		}
	}

	return macs
}

// router Attempt to get gateway IP, renamed from gateway to router due to package conflict
func router() net.IP {
	resp, err := gateway.DiscoverGateway()
	if err != nil {
		fmt.Println("[ERROR] Error getting Gateway: ", err.Error())
		log.Fatal(err)
	}
	return resp
}

// publicIP Attempt to get public IP address
func publicIP() net.IP {
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