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

	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		log.Fatal("Invalid address.")
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Printf("Address: %s\n", addr.String())
	fmt.Printf("Default mask: %v\n", bits)
	fmt.Printf("Leading ones count: %v\n", ones)
	fmt.Printf("Mask is (hex): %s\n", mask.String())
	fmt.Printf("Network is: %s\n", network.String())
}
