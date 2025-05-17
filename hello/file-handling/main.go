package main

import (
	"fmt"
	"os"
)

func main() {
	// // create a new file or truncate an existing one
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()

	// //Initial content to be added to the file
	// initialContent := "Hello, this is the initial content."

	// // write the initial content to the file using io.WriteString
	// _, errors := io.WriteString(file, initialContent+"\n")

	// if errors != nil {
	// 	fmt.Println("error while writing file:", errors)
	// 	return
	// }

	// fmt.Println("successfully created file")

	// Open the file
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("error while opening the file", err)
	// 	return
	// }
	// defer file.Close()

	// //create a buffer to read the file content
	// buffer := make([]byte, 1024)

	// //read the file content into the buffer.
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break // End of file reached
	// 	}

	// 	if err != nil {
	// 		fmt.Println("error reading file:", err)
	// 		return
	// 	}

	// 	// Proces the read content
	// 	fmt.Println(string(buffer[:n]))
	// }

	// instead you can use this
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	// Process the file content (in this example, just print it
	fmt.Println(string(content))
}
