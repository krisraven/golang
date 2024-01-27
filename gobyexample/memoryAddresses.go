// pointers
package main

import "fmt"

func main() {
	var a string = "hello world"
	fmt.Println(a) // "hello world"
	fmt.Println(&a) // 0xc000017400

	var b *string = &a // whatevr is assigned to b must be a point string
	fmt.Println(b) // 0xc000017400
	fmt.Println(*b) // "hello world"
	fmt.Println(&b) // 0xc000084280
}
