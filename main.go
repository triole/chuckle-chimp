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
	fmt.Println("\n")
	fmt.Printf(joke)
	fmt.Println("\n")

}
