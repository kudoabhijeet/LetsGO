package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Enter a Floating Point No.: ")
	fmt.Scan(&x)
	x = math.Trunc(x)
	fmt.Print("Truncated to :")
	fmt.Print(x)
}
