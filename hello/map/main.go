package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 50
	studentGrades["kartik"] = 60

	fmt.Println("grades of kartik:", studentGrades["kartik"])

	value, exists := studentGrades["david"]

	fmt.Println("grade of david is:", value)
	fmt.Println("david exists:", exists)

	for index, value := range studentGrades {
		fmt.Printf("key: %s and the value: %d\n", index, value)
	}

	//other way of declaring map

	student := map[string]int{
		"kartik": 90,
		"karan":  80,
		"pranav": 57,
	}

	for index, value := range student {
		fmt.Printf("key: %s and the value: %d\n", index, value)
	}

}
