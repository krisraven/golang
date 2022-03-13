package main

import "fmt"

func main() {
	// basic loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for loop w/o condition will loop forver and only stop if break or return
	for j := 2; j <= 10; j++  {
		fmt.Println("loop")
		break
	}

	// loop + skip item and continue to next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
	