package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s hostname", os.Args[0])
	}
	name := os.Args[1]
	addrs, err := net.LookupHost(name)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
}
