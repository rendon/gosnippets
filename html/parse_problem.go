// Parse Codeforces problem and extract sample test cases.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func parse(node *html.Node, mode, input, output *string) {
	var found = false
	var data = node.Data
	if node.Type == html.ElementNode && (data == "div" || data == "br") {
		if data == "div" {
			for _, a := range node.Attr {
				if a.Key == "class" && a.Val == "sample-test" {
					found = true
					*mode = "input"
					break
				}
			}
		} else if data == "br" {
			if *mode == "input" {
				*input += "\n"
			} else if *mode == "output" {
				*output += "\n"
			}
		}
	} else if node.Type == html.TextNode && *mode != "" {
		if data == "Input" {
			if *input != "" && *output != "" {
				fmt.Printf("Input:\n%s\n", *input)
				fmt.Printf("Output:\n%s\n", *output)
			}
			*input = ""
			*mode = "input"
		} else if data == "Output" {
			*output = ""
			*mode = "output"
		} else if *mode == "input" {
			*input += data
		} else if *mode == "output" {
			*output += data
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		parse(c, mode, input, output)
	}
	if found && (*input != "" || *output != "") {
		fmt.Printf("Input:\n%s\n", *input)
		fmt.Printf("Output:\n%s\n", *output)
	}
}

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatalf("USAGE: %s <problem_url>", os.Args[0])
	}

	var url = os.Args[1]
	var resp, err = http.Get(url)
	if err != nil {
		log.Fatalf("Error retrieving page: %s", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Error parsing document: %s", err)
	}
	var mode, input, output string
	parse(doc, &mode, &input, &output)
}
