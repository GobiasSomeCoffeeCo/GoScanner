package port

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

)

type State struct {
	PortsToScan int
	Service     string
	Protocol    string
	Address     string
	Threads     int
}


type ScanResult struct {
	Hostname State
	Port     int
	Status   string
}

// Determines if port is open or close based on net.DialTimeout
func ScanPorts(s *State, p int) ScanResult {
	result := ScanResult{Port: p, Hostname: *s}
	address := s.Address + ":" + strconv.Itoa(p)
	conn, err := net.DialTimeout(s.Protocol, address, 60*time.Second)
	if err != nil {
		result.Status = "Closed"
		return result
	}

	defer conn.Close()
	result.Status = "Open"

	return result
}

// Iterates through the ports specified and implements ScanPorts()
func InitialScan(s *State) {
	// Seperating protocols for now in case I want to handle them differently in the future
	if strings.ToLower(s.Protocol) == "tcp" {
		for i := 0; i <= s.PortsToScan; i++ {
			results := ScanPorts(s, i)
			fmt.Println(results)
		}
	} else if strings.ToLower(s.Protocol) == "udp" {
		for i := 0; i <= s.PortsToScan; i++ {
			results := ScanPorts(s, i)
			fmt.Println(results)
		}
	} else {
		fmt.Printf("Sorry, I do not understand this protocol: %s\n\r\n\r", s.Protocol)
	}
}


