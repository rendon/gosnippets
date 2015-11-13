// Compile C++ programs calling g++ compiler.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("USAGE: %s <src_file>", os.Args[0])
	}

	file := os.Args[1]
	re := regexp.MustCompile("\\.([^.]+)$")
	ext := re.FindString(file)
	out := os.TempDir() + "/" + strings.TrimSuffix(file, ext)
	cmd := exec.Command("g++", "-W", "-o", out, file)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Compiling %s\n", file)
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output file: %s\n", out)
}
