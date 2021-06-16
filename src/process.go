package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed data/jokes.json
var s []byte

// Joke struct is literally the funniest struct I've ever seen
type Joke struct {
	Joke string `json:"joke"`
}

func getJokes() (jokes []Joke) {
	var err error
	err = json.Unmarshal(s, &jokes)
	checkErr(err)
	return
}

func stringify(n Joke) string {
	return fmt.Sprintf("%s", n)
}
