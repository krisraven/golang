package main

import "fmt"

type rect struct {
	width, height int
}
// the area method has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined for either a pointer or value receiver types. This is an exmaple of a value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
// call 2 methods defined for our struct
	fmt.Println("area:	", r.area())
	fmt.Println("perim:	", r.perim())
// go automatically handles conversion between values and pointers for method calls
	rp := &r
	fmt.Println("area:	", rp.area())
	fmt.Println("perim:	", rp.perim())
}

