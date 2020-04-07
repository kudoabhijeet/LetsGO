package main

import "fmt"

func main() {
	var a int = 100
	var b int
	var ptr *int // Integer Pointer

	ptr = &a // Pointer 'ptr' now points to 'a'
	b = *ptr // Dereferencing to access data
	fmt.Printf(fmt.Sprintf("A -%d", a))
	fmt.Printf(fmt.Sprintf("B -%d", b))
}
