package main

import (
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// Connect to our memcache instance
	mc := memcache.New("127.0.0.1:11111")
	err := mc.Set(&memcache.Item{Key: "key_one", Value: []byte("michael")})
	if err != nil {
		log.Fatal(err)
	}
	err = mc.Set(&memcache.Item{Key: "key_two", Value: []byte("programming")})
	if err != nil {
		log.Fatal(err)
	}

	val, err := mc.Get("key_one")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("key_one: %s\n", val.Value)

	it, err := mc.GetMulti([]string{"key_one", "key_two"})
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range it {
		fmt.Printf("%s -> %s\n", k, v.Value)
	}
}
