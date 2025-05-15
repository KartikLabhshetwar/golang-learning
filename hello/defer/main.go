package main

import "fmt"

func main() {
	fmt.Println("start of the program")

	// the function inside the defer will be executed when the surrounding function ( main in this case ) exits
	defer fmt.Println("this will be executed at the end")

	fmt.Println("middle of the program")

	fmt.Println("start of the program")

	defer fmt.Println("This will be executed second")

	defer fmt.Println("This will be executed first")

	fmt.Println("middle of the program")
}
