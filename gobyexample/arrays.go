package main

import "fmt"

func main() {
	var a [5]int // create an array that will hold 5 ints. By default an array is 0-valued, which for int is 0
	fmt.Println("emp: ", a)


	a[4] = 100 // can set a value of an index. Eg 5th value is 100
	fmt.Println("set: ", a) // the full array
	fmt.Println("get: ", a[4]) // the actual value

	fmt.Println("len: ",len(a)) // built-in to return the length

	b :=[5]int{1, 2, 3, 4, 5} // declare and init array in one line
	fmt.Println("dcl: ", b)

	var twoD [2][3]int // arrays are 1-dimensional but multi-dimensional arrays can be built
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}