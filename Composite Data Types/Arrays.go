package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 2
	// fmt.Print(x[0])
	//Array Literals
	var y [5]int = [5]int{10, 20, 30}
	// ... is used to represent size of array
	// x: = [...]{1,3,4,5}

	// Iterating through Arrays
	for i, v := range y {
		fmt.Printf("Index : %d, Value : %d \n", i, v)
	}
}
