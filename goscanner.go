package main

import (
	"flag"
	"fmt"

	"github.com/GobiasSomeCoffeeCo/goscanner/port"
)

func main() {
	fmt.Println("\n\r============================================")
	fmt.Println("Starting up Goscanner...")
	fmt.Println("============================================\n\r")

	// var wg sync.WaitGroup
	// wg.Add(1)
	state := ParseCmdLine()
	// go func() {
	port.InitialScan(state)
	// 	wg.Done()
	// }()
	// wg.Wait()

}

func ParseCmdLine() *port.State {
	res := port.State{}

	valid := true

	flag.IntVar(&res.PortsToScan, "p", 1024, "Number of ports to scan")
	flag.StringVar(&res.Address, "u", "localhost", "The target URL or Domain")
	flag.StringVar(&res.Protocol, "pT", "tcp", "The protocol type you'd like to target (default tcp)")
	flag.IntVar(&res.Threads, "t", 4, "The amount of threads you'd like to work")

	flag.Parse()

	if res.PortsToScan <= 0 {
		fmt.Println("Ports (-p): Invalid Value: ", res.PortsToScan)
		valid = false
	}

	if res.Address == "" {
		fmt.Println("Url/Domain (-u): Must be specified")
		valid = false
	}

	if valid {
		return &res
	}

	return nil
}
