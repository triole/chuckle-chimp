package main

import (
	"encoding/json"
)

// Jokes a list of chuck norris jokes
// type JokesToDic struct {
// 	ID   int    `json:"id"`
// 	Joke string `json:"joke"`
// }

// Joke struct to get key from json
type Joke struct {
	Joke string `json:"joke"`
}

// GetAllJokes reads the jokes dictionary and returns a list of all jokes
func GetAllJokes() []Joke {
	var allJokes []Joke
	var tmpJokes []Joke
	input, err := Asset("data/jokes.json")
	checkErr(err)
	err = json.Unmarshal(input, &tmpJokes)
	checkErr(err)
	allJokes = append(allJokes, tmpJokes...)
	return allJokes
}
