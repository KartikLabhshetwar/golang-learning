package main

import "fmt"

func hello() {
	fmt.Println("hello there")
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("in the main function currently")
	hello()

	data := add(4, 8)
	fmt.Println("The sum of two numbers is:", data)

	ans := multiply(17, 45)
	fmt.Println("The multiplication of two number is:", ans)
}
