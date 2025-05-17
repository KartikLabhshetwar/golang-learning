package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning about web request.")

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error making the request", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("error while reading the response", err)
	}

	fmt.Println(string(data))

}
