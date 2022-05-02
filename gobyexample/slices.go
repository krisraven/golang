// slices are more powerful than arrays

package main 

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp: ", s) // [  ]


	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s) // [a b c]
	fmt.Println("set:", s[2]) // s

	fmt.Println("len: ", len(s))  // 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s) // [a b c d e f]

	c := make([]string, len(s))
	copy(c, s) // can copy a slice to a new slice
	fmt.Println("cpy: ", c) // [a b c d e f]

	slice1 := s[2:5]
	fmt.Println("sl1: ", slice1)// [c d e]

	slice2 := s[:5]
	fmt.Println("sl2: ", slice2) // [a b c d e]

	slice3 := s[2:]
	fmt.Println("sl3: ", slice3) // [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t) // [g h i]

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // [[0] [1 2] [2 3 4]]

	scores := make([]int, 0, 10)
// 	scores[7] = 9033 // cant do this because the array is has no data, so there is no index 7
    scores = scores[0:8] // need to set the whole array
    scores[7] = 9033 // then we can set a value
    scores[1] = 9027
	fmt.Println(scores)
}