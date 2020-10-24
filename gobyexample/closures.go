package main

import "fmt"
// this function intSeq returns another function which we define annoymously in the body of intSeq. The returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq() // intSeq is called assigning the result (a function) to nextInt. This function value captures its own i value which will be updated each time we call nextInt

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	// to confirm that the state is unique to a particular function we create a test a new one
	newInts := intSeq()
	fmt.Println(newInts()) // 1
}