package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <hostname>", os.Args[0])
	}

	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		log.Fatalf("Failed to resolve name: %s", err)
	}
	fmt.Printf("IP of %s: %s\n", name, addr.String())
}
