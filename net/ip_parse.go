// Source: Network programming with Go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <ip>", os.Args[0])
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Printf("Invalid address.\n")
	} else {
		fmt.Printf("The address is %s\n", addr.String())
	}
}
