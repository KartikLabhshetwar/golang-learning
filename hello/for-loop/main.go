package main

import "fmt"

func main() {
	//eg 1

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	counter := 0

	for {
		fmt.Println("infinite loop")
		counter++
		if counter == 3 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5}

	for i, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	message := "hello world"

	for index, char := range message {
		fmt.Printf("Index: %d, Value: %c\n", index, char)
	}

}
