package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(s1)
	fmt.Println(s2)
	// Length and Capacity
	fmt.Println(len(s1), cap(s1))
	// Accessing Slices
	fmt.Println(s1[1])
	fmt.Println(s2[0])
	//Slice Literals
	sli:= []int{1,2,3,4,5}
	fmt.Println(sli)
}
