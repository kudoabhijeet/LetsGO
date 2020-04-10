package main

import "fmt"

func main() {
	// float32 6 digits of precision
	// float64 15 digits of precision
	var x float64 = 123.45           // basic representation
	var y float64 = 1.234e2          // scientific representation
	var z complex128 = complex(1, 3) // complex no. 1 - real part, 3- imaginary part (1+3i)
	fmt.Print(x)
	fmt.Print(y)
	fmt.Print(z)

}
