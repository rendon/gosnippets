package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	fmt.Printf("%x", sha1.Sum([]byte("abc")))
}
