package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8, 9)

	fmt.Println("slice:", numbers)
	fmt.Println(len(numbers))

	stock := make([]int, 3, 5)

	fmt.Println("slice:", stock)
	fmt.Println("length:", len(stock))
	fmt.Println("capacity:", cap(stock))
}
