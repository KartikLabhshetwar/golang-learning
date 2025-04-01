package main

import (
	"fmt"
	"hello/util"
)

func main() {
	fmt.Println("this is the first program in go")

	var hello string = "hello"
	var digit int = 23
	var isTrue bool = false

	newVar := "hello"

	fmt.Println(hello)
	fmt.Println(digit)
	fmt.Println(isTrue)
	fmt.Println(newVar)

	util.PrintMessage("this is the first custom package in go")
}
