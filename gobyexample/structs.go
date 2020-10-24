package main

import "fmt"

type person struct {
	name 	string
	age		int 
}

func newPerson(name string) *person {

	p := person{name: name} // can safely return a pointer to a local variable as a local variable will survive the scope of the function
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20}) // this creates a new struct
	fmt.Println(person{name: "Alice", age: 20})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40}) // an & prefix yields a pointer to the struct
	fmt.Println(newPerson("Jon")) // its idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // access sturct fields with a dot

	sp := &s // can use dot with struct pointers
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age) // structs are mutable
}