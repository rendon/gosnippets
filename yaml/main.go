// Parse and display YAML document.
// You'll need a file called `sample.yml` with some YAML data next to this file,
// e.g.:
//     ---
//     key: value
//     A:
//         AA: value
//         AB: value
//         AC: value
//     B:
//         BA: value
//         BB: value
//         BC: value
//     C:
//         CA: value
//         CB: value
//         CC: value

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func traverse(node interface{}, tabs string) error {
	var data, ok = node.(map[interface{}]interface{})
	if !ok {
		errors.New("Failed to determine node type.")
	}

	for k, v := range data {
		var value, ok = v.(string)
		fmt.Printf("%s%s:", tabs, k)
		if ok {
			fmt.Printf(" %s\n", value)
		} else {
			fmt.Println()
			if err := traverse(v, tabs+"    "); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	log.SetFlags(0)
	var buffer []byte
	var err error
	buffer, err = ioutil.ReadFile("sample.yml")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var data map[interface{}]interface{}
	if err = yaml.Unmarshal(buffer, &data); err != nil {
		log.Fatalf("Failed to unmarshal content: %s", err)
	}

	err = traverse(data, "")
	if err != nil {
		log.Fatalf("Failed to traverse data: %s", err)
	}
}
