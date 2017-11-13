package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// --- utility functions
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// CleanUpString removes unnecessary characters from string
func CleanUpString(s string) string {
	r, _ := regexp.Compile("[^{}]+")
	arr := r.FindAllString(s, -1)
	rs := strings.Join(arr[:], ",")
	return rs
}

// PrintFile prints file content to stdout
func PrintFile(fn string) {
	data, err := ioutil.ReadFile(fn)
	checkErr(err)
	fmt.Println(string(data))
}

// RandomNumberBetween returns a random number in a range between two values
func RandomNumberBetween(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// WordWrap wraps a string making it fit a line length, respects word boundaries
func WordWrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped

}
