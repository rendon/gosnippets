package main

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var reDate = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d$`)
var reDateTime = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d$`)

var loc *time.Location

func init() {
	var err error
	loc, err = time.LoadLocation("UTC")
	if err != nil {
		log.Fatalf("Failed to load location: %s", err)
	}
}

// parseDate this function assumes str matches 'YYYY-MM-DD'.
func parseDate(str string) (int, int, int) {
	tokens := strings.Split(str, "-")
	year, _ := strconv.Atoi(tokens[0])
	month, _ := strconv.Atoi(tokens[1])
	day, _ := strconv.Atoi(tokens[2])
	return year, month, day
}

func ParseTime(str string) (time.Time, error) {
	str = strings.Trim(str, " ")
	if !reDate.MatchString(str) && !reDateTime.MatchString(str) {
		return time.Now(), errors.New("No valid date found")
	}
	if reDate.MatchString(str) {
		year, month, day := parseDate(str)
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc), nil
	}
	tokens := strings.Split(str, " ")
	year, month, day := parseDate(tokens[0])

	t := strings.Split(tokens[1], ":")
	// Ignore errors since we know the input is valid
	hour, _ := strconv.Atoi(t[0])
	min, _ := strconv.Atoi(t[1])
	sec, _ := strconv.Atoi(t[2])
	return time.Date(year, time.Month(month), day, hour, min, sec, 0, loc), nil
}

func main() {
	if !reDate.MatchString("2016-07-19") {
		log.Fatal("Expected to match valid date")
	}
	if reDate.MatchString("2016-07/19") {
		log.Fatal("It should not match invalid date")
	}
	if reDate.MatchString("2016-mm-19") {
		log.Fatal("It should not match invalid date")
	}
	if reDate.MatchString("2016-07-19 22:45:31") {
		log.Fatal("It should not match datetime")
	}

	if !reDateTime.MatchString("2016-07-19 22:45:31") {
		log.Fatalf("Expected to mathc valid datetime")
	}

	if reDateTime.MatchString("2016-07-1922:45:31") {
		log.Fatal("It should not match invalid datetime")
	}
	if reDateTime.MatchString("2016-07-19 22/45/31") {
		log.Fatal("It should not match invalid datetime")
	}
	if reDateTime.MatchString(" 2016-07-19 22:45:31") {
		log.Fatal("It should not match invalid datetime")
	}

	t, _ := ParseTime("2016-07-19")
	if t.Year() != 2016 || t.Month() != 7 || t.Day() != 19 {
		log.Fatal("Expected to parse valid date")
	}

	t, _ = ParseTime("2016-07-19 18:55:03")
	if t.Year() != 2016 || t.Month() != 7 || t.Day() != 19 ||
		t.Hour() != 18 || t.Minute() != 55 || t.Second() != 3 {
		log.Fatal("Expected to parse valid datetime")
	}
}
