package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

var (
	jokesList     []string
	jokesJSONFile = "/home/ole/rolling/golang/projects/chuckjoke/jokes.json"
)

// Jokes a list of chuck norris jokes
// type JokesToDic struct {
// 	ID   int    `json:"id"`
// 	Joke string `json:"joke"`
// }

// JokesToList create a simple array containing all the jokes from the json
type JokesToList struct {
	Joke string `json:"joke"`
}

func (t JokesToList) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func getJokes() []JokesToList {
	jokes := make([]JokesToList, 1000000)
	raw, err := ioutil.ReadFile(jokesJSONFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(raw, &jokes)
	return jokes
}

func main() {
	jokes := getJokes()
	for _, te := range jokes {
		// jokesList = append(jokesList, te)
		fmt.Println(te)
	}
}

// --- utility functions
func shuffleList(a []string) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
