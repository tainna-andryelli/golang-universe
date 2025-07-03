package main //declare a main package

import (
	"fmt" // import the fmt package which contains functions for formatting text

	"rsc.io/quote"
)

// A main function executes by default when you run the main package
func main() {
	fmt.Println(quote.Go()) // Print "Hello, world!" to the console
}
