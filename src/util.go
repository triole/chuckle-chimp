package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func cleanUpString(s string) string {
	r, _ := regexp.Compile("[^{}]+")
	arr := r.FindAllString(s, -1)
	rs := strings.Join(arr[:], ",")
	return rs
}

func printFile(fn string) {
	data, err := ioutil.ReadFile(fn)
	checkErr(err)
	fmt.Println(string(data))
}

func randomNumberBetween(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func wordWrap(text string, lineWidth int) string {
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
