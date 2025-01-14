package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	device := "wlan0"

	log.Default()
	log.SetFlags(0)

	handle, err := pcap.OpenLive(device, 1600, true, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packet := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println(packet)

	for packets := range packet.Packets() {
		ethernetLayer := packets.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			fmt.Printf("Ethernet Frame: %s -> %s\n", ethernetPacket.SrcMAC, ethernetPacket.DstMAC)
		}

		ipLayer := packets.Layer(layers.LayerTypeIPv4)
		if ipLayer == nil {
			ipLayer = packets.Layer(layers.LayerTypeIPv6)
		}

		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			fmt.Printf("IP Packet: %s -> %s\n", ip.SrcIP, ip.DstIP)
		}

		tcpLayer := packets.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			fmt.Printf("TCP Segment: %d -> %d\n", tcp.SrcPort, tcp.DstPort)
		}

	}

}
