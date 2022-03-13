package main

import "fmt"

// this function has an int parameter so arguments will be passedd to it by value
func zeroval(ival int) {
	ival = 0 // when this function is called ival will be different
}

// this function has a *int parameter meaning that it takes an int pointer. The *iptr dereferences the pointer from its memory address to the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 100
	j := 200
	k := 300
	l := 400
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	// the &i syntax gives the memory address of a pointer to i
	fmt.Println("pointer:", &i)
	fmt.Println("pointer:", &j)
	fmt.Println("pointer:", &k)
	fmt.Println("pointer:", &l)
}