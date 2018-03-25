package main

import (
	"fmt"
)

func main() {
	jokes := GetAllJokes()
	r := RandomNumberBetween(0, len(jokes))
	joke := StringifyJoke(jokes[r])
	joke = CleanUpString(joke)

	wrapped := WordWrap(joke, 30)
	fmt.Println("")
	fmt.Println("")
	fmt.Printf(wrapped)
	fmt.Println("")
	fmt.Println("")
}
