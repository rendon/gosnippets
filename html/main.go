// Parse sample tests of Codeforces problem.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

var input string
var output string

func traverse(node *html.Node, depth int, mode *string) bool {
	if node.Type == html.ElementNode && (node.Data == "div" || node.Data == "br") {
		if node.Data == "div" {
			var found = false
			for _, a := range node.Attr {
				if a.Key == "class" && a.Val == "sample-test" {
					found = true
					break
				}
			}
			if found {
				depth = 1
			}
		} else if node.Data == "br" {
			if *mode == "input" {
				input += "\n"
			} else if *mode == "output" {
				output += "\n"
			}
		}
	} else if node.Type == html.TextNode && depth > 0 {
		//fmt.Printf("%q\n", node.Data)
		if node.Data == "Input" {
			if input != "" && output != "" {
				fmt.Printf("Input:\n%s\n", input)
				fmt.Printf("Output:\n%s\n", output)
			}
			input = ""
			*mode = "input"
		} else if node.Data == "Output" {
			output = ""
			*mode = "output"
		} else if *mode == "input" {
			input += node.Data
		} else if *mode == "output" {
			output += node.Data
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if depth == 2 {
			fmt.Printf("Reading input...\n")
		}
		traverse(c, depth, mode)
	}

	return false
}

func main() {
	var resp, err = http.Get("http://codeforces.com/contest/479/problem/A")
	if err != nil {
		log.Fatalf("Error retrieving page: %s", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Error parsing document: %s", err)
	}
	var mode = ""
	traverse(doc, 0, &mode)
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", output)
}
