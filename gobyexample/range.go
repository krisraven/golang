package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// we don't need the index so we can use the _ identifier
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum) // 9

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i) // 1
		}
	}

	kvs := map[string]string{"a": "apple", "d": "dragon"} // range on map iterates of key/value pairs
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // a -> apple   b -> banana
	}
	// range can just iterate over keys of a map
	for k := range kvs {
		fmt.Println("key: ", k) // b   a
	}
	// range on strings iterates of Unicode code points
	for i, c := range "recipe" {
		fmt.Println(i, c) // 0 114	1 101	2 99	3 105	4 112	5 101
	}
}