package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("learning about JSON.")
	person := Person{Name: "kartik", Age: 21, IsAdult: true}
	fmt.Println("The data of the person is:", person)

	//convert the person into JSON encoding (marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error while marshalling", err)
		return
	}

	fmt.Println("person data is:", string(jsonData))

	// decoding (unmarshalling)

	var personData Person

	err = json.Unmarshal(jsonData, &personData)

	if err != nil {
		fmt.Println("error while unmarshalling the data", err)
		return
	}

	fmt.Println("Person data after unmarshalling is:", personData)
}
