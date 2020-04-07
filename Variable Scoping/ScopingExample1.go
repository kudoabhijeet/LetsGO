package main

import "fmt"

func main() {
	first()
	second()
}

func first() {
	var x int = 100
	fmt.Printf("%d", x) // Prints the value of 'x'
}
func second() {
	fmt.Printf("%d", x) // Prints the value of 'x'
}

// On calling the 'second()' function the compiler returns an error : 'Undefined x'
