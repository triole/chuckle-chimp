package main

import (
	"fmt"
)

func main() {
	jokes := getJokes()
	r := randomNumberBetween(0, len(jokes))
	joke := stringify(jokes[r])
	joke = cleanUpString(joke)

	wrapped := wordWrap(joke, 50)
	fmt.Printf("\n\n%s\n\n", wrapped)
}
