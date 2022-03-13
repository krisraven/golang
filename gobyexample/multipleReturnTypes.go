package main

import "fmt"
// the int, int in this function signature shows that the function returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	c, _ := vals() // if a subset of the returned values is needed use the _ identifier
	// there is 2 values. 1st is chosen (ie a = c) and 2nd one is blank (ie _ = a)
	fmt.Println(c)
}