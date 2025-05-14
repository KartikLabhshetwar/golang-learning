package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "apple, banana, orange"
	newStr := strings.Split(str, " ")
	fmt.Println(newStr)

	str1 := "one two three two two three one"
	occ := strings.Count(str1, "two")
	fmt.Println(occ)

	str3 := "    hello, go!    "
	trimmed := strings.TrimSpace(str3)
	fmt.Println(trimmed)

	str4 := "Kartik"
	str5 := "Labhshetwar"
	conc := strings.Join([]string{str4, str5}, " ")
	fmt.Println(conc)

}
