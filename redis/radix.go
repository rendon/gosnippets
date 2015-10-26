package main

import (
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("Couldn't connect to Redis server")
	}
	defer client.Close()

	val, err := client.Cmd("ZCARD", "myxyht").Int64()
	if err != nil {
		log.Fatalf("Error retrieving value: %s", err)
	}
	fmt.Printf("value = %d", val)
}
