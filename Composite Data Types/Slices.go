package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(s1)
	fmt.Println(s2)
	// Length and Capacity
	fmt.Print(len(s1), cap(s1))

}
