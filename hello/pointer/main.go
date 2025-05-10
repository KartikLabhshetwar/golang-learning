package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 4
}

func main() {
	var num int
	num = 10

	var ptr *int

	ptr = &num

	// a more simple way is this

	data := 20
	p := &data
	fmt.Println("The value is:", data)
	fmt.Println("the address of num is:", p)
	fmt.Println("the value of num through address ptr is:", *p)

	fmt.Println("The value is:", num)
	fmt.Println("the address of num is:", ptr)
	fmt.Println("the value of num through address ptr is:", *ptr)

	var empty_pointer *int

	if empty_pointer == nil {
		fmt.Println("Pointer is nil")
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Modified value:", value)
}
