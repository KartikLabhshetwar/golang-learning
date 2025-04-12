package main

import "fmt"

func main() {
	x := 10

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x > 5 {
		fmt.Println("x is greater than 5 but less than 10")
	} else {
		fmt.Println("x is less than 5")
	}

	y := 2

	if y > 2 && x > 10 {
		fmt.Println("hello world")
	} else {
		fmt.Println("hi there!")
	}
}
