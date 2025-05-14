package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 20

	var floatNum float64 = float64(num)
	fmt.Println("converson from integer to float", floatNum)

	number := 2

	str := strconv.Itoa(number)
	fmt.Println("converson from integer to string", str)

	strNum := "23232"

	numbers, _ := strconv.Atoi(strNum)
	fmt.Println("converson from string to integer", numbers)

	str1 := "3.222"

	n, _ := strconv.ParseFloat(str1, 64)
	fmt.Println("converson from string to float", n)

}
