package main

import "fmt"

func main() {
	var fruits [5]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fruits[3] = "Mango"
	fruits[4] = "Grapes"

	fmt.Println(fruits)
	fmt.Println(fruits[2])

	//for getting the length of the array
	fmt.Println(len(fruits))
}
