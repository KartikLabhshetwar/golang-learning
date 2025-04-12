package main

import (
	"fmt"
)

func main() {
	day := 2

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	default:
		fmt.Println("no day exist here")
	}

	month := "jan"

	switch month {
	case "jan", "feb", "march":
		fmt.Println("winter")
	case "april", "may":
		fmt.Println("summer")
	default:
		fmt.Println("nothing")
	}

	temp := 25

	switch {
	case temp < 0:
		fmt.Println("cold")
	case temp > 0 && temp < 20:
		fmt.Println("cool")
	case temp > 20 && temp < 30:
		fmt.Println("warm")
	default:
		fmt.Println("hot")
	}

}
