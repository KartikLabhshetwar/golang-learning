package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning url")

	myURL := "https://example.com:8080/path/to/resources?key1=value1&key2=value2"
	fmt.Printf("type of URL %T\n", myURL)

	parseURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("error parsing url", err)
		return
	}

	fmt.Printf("type of url: %T\n", parseURL)

	fmt.Println("scheme", parseURL.Scheme)
	fmt.Println("host", parseURL.Host)
	fmt.Println("path", parseURL.Path)
	fmt.Println("RawQuery", parseURL.RawQuery)

	// modifiying the URL components

	parseURL.Path = "/newPath"
	parseURL.RawQuery = "username=iamkartik"

	//constructing a url string from a url object
	newURL := parseURL.String()
	fmt.Println(newURL)
}
