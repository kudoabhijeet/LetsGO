package main

import "fmt"

func main() {
	//var idMap map[string]int
	//idMap = make(map[string]int) // creates a empty map
	//literal
	id := map[string]int{"joe": 123}
	//Accessing value
	fmt.Print(id["joe"])
	// Updating/ assigning
	id["jane"] = 456
	//deleting a pair
	delete(id, "joe")
	//iterating through map
	for key, val := range id {
		fmt.Println(key, val)
	}
}
