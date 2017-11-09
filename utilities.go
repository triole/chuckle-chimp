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

func shuffleList(a []Joke) []Joke {
	for i := range a {
		j := RandomNumberBetween(0, 10)
		a[i], a[j] = a[j], a[i]
	}
	return a
}
