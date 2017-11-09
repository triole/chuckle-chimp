package main

import (
	"math/rand"
	"time"
)

// --- utility functions
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
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
