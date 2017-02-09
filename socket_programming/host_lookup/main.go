package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <domain>\n", os.Args[0])
		os.Exit(1)
	}
	host := os.Args[1]
	ips, err := net.LookupHost(host)
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	os.Exit(0)
}
