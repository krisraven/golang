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

	_, c := vals() // if a subsete of the returned values is needed use the _ identifier
	fmt.Println(c)
}