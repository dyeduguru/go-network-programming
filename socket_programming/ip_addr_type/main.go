package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <IP-Address>\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	ip := net.ParseIP(name)
	if ip == nil {
		fmt.Print("Invalid ip address\n")
	} else {
		fmt.Printf("The address is: %s\n", ip.String())
	}
}