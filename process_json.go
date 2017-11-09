package main

import "encoding/json"

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

	// two ways of accessing the json data
	// simple json file reading
	// raw, err := ioutil.ReadFile(jokesJSONFile)
	// using assets of go-bindata
	raw, err := Asset("data/jokes.json")

	// error handler, but you know that
	if err != nil {
		panic(err)
	}

	json.Unmarshal(raw, &jokes)
	return jokes
}
