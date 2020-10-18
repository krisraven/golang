package main

import "fmt"

func main() {
	// basic if statement
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// if statement without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	// braces are required, but not parentheses
	// a statement can precede conditionals. Variables are avaiable in all branches
	if num := 11; num < 0 {
		fmt.Println(num, "is negtive")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
/// there's no ternary in go so a full if statement needs to be used