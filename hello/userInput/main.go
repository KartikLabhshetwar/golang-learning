package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string

	// fmt.Println("Enter your name:")
	// fmt.Scan(&name)

	// fmt.Printf("Hello %s!", name)

	fmt.Println("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)

}
