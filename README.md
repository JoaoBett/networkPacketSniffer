# Network Packet Sniffer in Go

A lightweight and efficient network packet sniffer written in Go. This tool captures and analyzes network packets in real time, providing insights into network traffic for debugging, monitoring, or educational purposes.

## Features

- Capture network packets in real time.
- Analyze packet details (e.g., source/destination IP, ports, protocols).
- Filter packets by protocol (TCP, UDP, ICMP, etc.).
- Save captured packets to a file for offline analysis.
- Cross-platform support (Linux, macOS, Windows).

## Requirements

- Go 1.20+ (or the latest stable version)
- libpcap/WinPcap (optional for advanced features on certain systems)
