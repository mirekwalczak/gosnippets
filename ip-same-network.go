package main

import (
	"fmt"
	"log"
	"net"
)

// isIpInTheSameNet compares two IP addresses
// and say if they are in the same subnet
func isIPInTheSameNet(ip1, ip2 string) bool {
	ipIP1, ipNet1, err := net.ParseCIDR(ip1)
	if err != nil {
		log.Fatal(err)
	}
	ipIP2, ipNet2, err := net.ParseCIDR(ip2)
	if err != nil {
		log.Fatal(err)
	}
	result := intersect1(ipNet1, ipNet2)
	fmt.Println("IP1 =", ipIP1, "; IP2 =", ipIP2)
	fmt.Println("Does", ipNet1, "contain", ipNet2, ":", result)
	fmt.Println()
	return result
}

func intersect1(ip1, ip2 *net.IPNet) bool {
	return ip2.Contains(ip1.IP) || ip1.Contains(ip2.IP)
}

func intersect2(n1, n2 *net.IPNet) bool {
	for i := range n1.IP {
		if n1.IP[i]&n1.Mask[i] != n2.IP[i]&n2.Mask[i]&n1.Mask[i] {
			return false
		}
	}
	return true
}

func main() {
	isIPInTheSameNet("10.111.164.0/26", "10.111.164.1/26")
	isIPInTheSameNet("10.0.0.0/8", "10.111.164.1/26")
	isIPInTheSameNet("1.1.0.2/16", "1.1.1.1/24")
	isIPInTheSameNet("1.1.1.1/24", "1.1.0.2/16")
	isIPInTheSameNet("10.0.0.0/24", "10.0.0.0/23")
}
