package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("Error finding devices: %v", err)
	}
	fmt.Println("Available devices:")
	for _, device := range devices {
		fmt.Printf("Name: %s, Description: %s\n", device.Name, device.Description)
		for _, addr := range device.Addresses {
			fmt.Printf("  IP Address: %s\n", addr.IP)
		}
	}
}
