package main

import "fmt"

var (
	jokesList     []string
	jokesJSONFile = "/home/ole/rolling/golang/projects/chuckjoke/jokes.json"
)

func main() {
	jokes := GetAllJokes()
	r := RandomNumberBetween(0, len(jokes))
	joke := jokes[r]
	fmt.Println(joke)
	// i := random(0, 90)
	// fmt.Println(i)
	// for _, file := range j {
	// fmt.Println(file)
	// }
	// fmt.Println(jokes[i])
}
