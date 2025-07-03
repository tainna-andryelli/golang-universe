package main

import "fmt"

var musicFavorite = "Emptiness Machine, Linkin Park" // var declares a package-level variable
// music := "Emptiness Machine, Linkin Park" // this will not work

func main() {
	var name = "Tainna" // will infer the type as string
	var age int = 22
	var cats, dogs = 2, 3 // var declares multiple variables

	fmt.Println(name, age, cats, dogs)

	music := "Emptiness Machine, Linkin Park" // short variable declaration only inside functions
	fmt.Println(music, musicFavorite)
}
