package main

import "fmt"


var t = "Variável acessível no pacote"

func main() {
	var v = "Variável acessível na função"

	bytesNumber, errors := fmt.Println("")

	fmt.Println(bytesNumber, errors)

	x := 10
	y := "Sitrus"

	x = 20 // novo valor

	x, z := 20, 30 // atribuição ao x, declaração ao z

	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("y: %v %T\n", y, y)
	fmt.Printf("z: %v %T\n", z, z)

	fmt.Println(t, "\n", v)
}