package main

import (
	"fmt"
)

var (
	jokesList     []string
	jokesJSONFile = "/home/ole/rolling/golang/projects/chuckjoke/jokes.json"
)

func main() {
	jokes := getJokes()
	for _, te := range jokes {
		jokesList = append(jokesList, te)
		fmt.Println(te)
	}
}
