package main

import (
	"fmt"
	"math"
)
// basic interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}
// to implement an interface in go we just need to implement all the methods in the interface. Here we implement "geometry" on "rect"
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implement "geometry" on "circle"
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
// If a variable has an interface type then we can call methods that are in the named interface. The "measure" function takes advantage of this to work on any "geometry"
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}