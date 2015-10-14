// Basic example using gosexy/redis (https://github.com/gosexy/redis) client.
package main

import (
	"fmt"
	"log"
	"menteslibres.net/gosexy/redis"
)

func main() {
	var rc = redis.New()
	var err error
	if err = rc.Connect("localhost", 6379); err != nil {
		log.Fatalf("Couldn't connect to redis server.")
	}

	if _, err = rc.HSet("key", "field", "value"); err != nil {
		log.Fatalf("Error setting value: %s", err)
	}

	var value string
	if value, err = rc.HGet("key", "field"); err != nil {
		log.Fatalf("Error getting value: %s", err)
	}

	fmt.Printf("Value = %s\n", value)
}
