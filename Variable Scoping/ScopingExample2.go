package main
import "fmt"
var x int = 100 // Declared and initialized outside function : main(), first(), second()

func main() {
	first()
	second()
}

func first (){
	
	fmt.Printf("%d\n",x) // Prints the value of 'x'
}
func second (){
	fmt.Printf("%d\n",x) // Prints the value of 'x'
}
