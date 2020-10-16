package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // variables declared with a init value are zero-valued
	fmt.Println(e)

	f:= "apple" // the := syntax is shorthand for declaring and intialising a variable
	fmt.Println(f)
}