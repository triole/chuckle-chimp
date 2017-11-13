package main

import (
	"fmt"
)

func main() {
	jokes := GetAllJokes()
	r := RandomNumberBetween(0, len(jokes))
	joke := StringifyJoke(jokes[r])
	joke = CleanUpString(joke)

	// final output
	wrapped := WordWrap(joke, 30)
	fmt.Println("\n")
	fmt.Printf(wrapped)
	fmt.Println("\n")

}
