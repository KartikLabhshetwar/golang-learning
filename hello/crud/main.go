package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performPostRequest() {
	todo := Todo{
		UserId:    22,
		Title:     "kartik labhshetwar",
		Completed: true,
	}

	//convert todo struct into JSON.
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error marshalling the data", err)
		return
	}

	// convert json data to string.
	jsonString := string(jsonData)

	// convert string to Reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// send POST request

	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error sending request", err)
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error while reading the response", err)
		return
	}

	fmt.Print(string(data))

	fmt.Print(res.StatusCode)

}

func performUpdateRequest() {
	todo := Todo{
		UserId:    2323,
		Title:     "hello world, I am here",
		Completed: false,
	}

	// convert struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error while marshalling")
	}

	// convert JSON to string
	jsonString := string(jsonData)

	// convert string to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Update method

	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("error doing put request", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending put request", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response:", err)
		return
	}
	fmt.Println(string(data))
	fmt.Print(res.StatusCode)

}

func performGetRequest() {
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

func perfromDeleteRequest() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Delete method

	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("error while requesting", err)
		return
	}

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending put request", err)
		return
	}
	defer res.Body.Close()

	fmt.Println(res.Status)

}

func main() {
	fmt.Println("learning about crud operations in go.")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	perfromDeleteRequest()
}
