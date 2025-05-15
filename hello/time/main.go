package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02")
	fmt.Println(formattedTime)

	dateStr := "2004-04-17"
	parseDate, _ := time.Parse("2006-01-02", dateStr)
	fmt.Println(parseDate)
}
