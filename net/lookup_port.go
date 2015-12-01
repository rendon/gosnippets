package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage %s network-type service", os.Args[0])
	}
	networkType := os.Args[1]
	service := os.Args[2]
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Service port %d", port)
}
