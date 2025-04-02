package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}

	return a / b, nil
}

func main() {

	ans, _ := divide(10, 4)

	fmt.Println("leanring about error handling")

	fmt.Printf("The division of two number is: %.3f\n", ans)
}
