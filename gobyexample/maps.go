package main

import "fmt"

func main() {
	m := make(map[string]int) // to create an empty map use make

	m["k1"] = 7 // set key/value pairs using this syntax
	m["k2"] = 13

	fmt.Println("map: ", m) // map[k1:7 k2:13]. The word map appears in the result when fmt.Println is used

	v1 :=m["k1"]
	fmt.Println("v1: ", v1) // 7
	fmt.Println("v1: ", len(m)) // 2

	delete(m, "k2") // removes key/value pairs from the map
	fmt.Println("map: ", m) // map[k1:7]

	_, prs := m["k2"] // optional 2nd return value indicates in the key was present in the map. This helps disambiguate between missing keys or keys with "" and 0 value. Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	fmt.Println("prs: ", prs) // false


	n := map[string]int{"foo": 1, "bar": 2} // declare and init a new map
	fmt.Println("map: ", n) // map[bar:2 foo:1]
}