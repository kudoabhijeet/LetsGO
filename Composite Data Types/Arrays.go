package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 2
	fmt.Print(x[0])
	//Array Literals
	var y [5]int = [5]{2,3,4,6,7}
	// ... is used to represent size of array
	// x: = [...]{1,3,4,5}
}
