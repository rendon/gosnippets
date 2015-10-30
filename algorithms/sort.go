package main

import (
	"fmt"
	"sort"
)

func main() {
	var items = []string{"z", "p", "a"}
	sort.Strings(items)
	for i, item := range items {
		fmt.Printf("%d: %s\n", i, item)
	}
}
