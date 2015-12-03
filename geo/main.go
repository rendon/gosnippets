// Detect coordinates and obtain address using Google Maps Geocoding API.

// Input file sample:
// Grand Rapids MI
// ÜT: 51.67985,-0.394467
// iPhone: 1.307380,103.884697
// Bay Area, CA

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kellydunn/golang-geo"
)

var (
	// Simple floating point regex
	r  = `[-+]?[0-9]+(.[0-9]*)?`
	re = regexp.MustCompile(fmt.Sprintf(`%s\s*,\s*%s`, r, r))
)

func parseCoordinate(c string) (*geo.Point, error) {
	c = strings.Replace(c, " ", "", -1)
	points := strings.Split(c, ",")
	if len(points) != 2 {
		return nil, errors.New("Invalid coordinates")
	}
	lat, err := strconv.ParseFloat(points[0], 64)
	if err != nil {
		return nil, err
	}
	long, err := strconv.ParseFloat(points[1], 64)
	if err != nil {
		return nil, err
	}
	return geo.NewPoint(lat, long), nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}
	fn := os.Args[1]
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	// You'll need a Google API key.
	geo.SetGoogleAPIKey(os.Getenv("GOOGLE_GEO_API_KEY"))
	geocoder := geo.GoogleGeocoder{}
	for _, line := range strings.Split(string(buf), "\n") {
		match := re.FindString(line)
		if match != "" {
			p, err := parseCoordinate(match)
			if err != nil {
				log.Println(err)
				continue
			}
			addr, err := geocoder.ReverseGeocode(p)
			if err != nil {
				log.Println(err)
				continue
			}
			tokens := strings.Split(addr, ",")
			log.Printf("Coordinates: %s\n", match)
			log.Printf("Address: %s\n", addr)
			log.Printf("Country: %s\n\n", tokens[len(tokens)-1])
		}
	}
}
