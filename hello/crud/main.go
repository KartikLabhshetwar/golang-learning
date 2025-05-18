package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("learning about crud operations in go.")

	//learning about get request in go

	req, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error while making the request", err)
		return
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		fmt.Print("error in getting response", req.Status)
		return
	}

	// data, err := ioutil.ReadAll(req.Body)

	// if err != nil {
	// 	fmt.Println("error reading", err)
	// 	return
	// }

	// fmt.Println("the get request is:", string(data))

	var todo Todo

	err = json.NewDecoder(req.Body).Decode(&todo)

	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Print("Todo:", todo)

}
