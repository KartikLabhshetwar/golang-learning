package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello!")
}

func sayHi() {
	fmt.Println("Hi kartik!..")
}

func main() {

	//start a new goroutine
	go sayHello()
	sayHi()

	// wait for a moment to allow the goroutines to finish
	time.Sleep(1000 * time.Millisecond)

}
