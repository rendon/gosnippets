package main

import (
	"log"
	"menteslibres.net/gosexy/redis"
)

func main() {
	var rc = redis.New()
	var err error
	if err = rc.Connect("redis-server", 6379); err != nil {
		log.Fatalf("Couldn't connect to redis server.")
	}

	var value string
	if value, err = rc.HGet("key", "field"); err != nil {
		log.Fatalf("Error getting value: %s", err)
	}

	log.Printf("Value = %s", value)
}
