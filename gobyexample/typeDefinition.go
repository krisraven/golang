package main

import(
	"fmt"
)

type TD struct {
	Foo string
}

func (td *TD) Bar() { // *TD is as type definition
	 td.Foo = `bar`
 }

 func main() {
	 a := new(TD)
	 a.Bar()
	 fmt.Println(a.Foo) // bar
 }