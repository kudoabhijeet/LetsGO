package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter a String :")
	var x string
	fmt.Scanln(&x)
	x = strings.ToLower(x)
	// fmt.Print(x)
	if strings.Contains(x, "a") && strings.HasPrefix(x, "i") && strings.HasSuffix(x, "n") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}

}
