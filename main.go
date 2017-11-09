package main

import (
	"fmt"
)

func main() {
	jokes := GetAllJokes()
	r := RandomNumberBetween(0, len(jokes))
	joke := StringifyJoke(jokes[r])
	joke = CleanUpString(joke)
	fmt.Println("\n")
	fmt.Printf(joke)
	fmt.Println("\n")

	// i := random(0, 90)
	// fmt.Println(i)
	// for _, file := range j {
	// fmt.Println(file)
	// }
	// fmt.Println(jokes[i])
}
