package main

import (
	"fmt"
)

func welcome(text []string) {

	fmt.Printf("%s %s", text0, text1)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	inputString := []string{
		"Welcome to 30 Days of Code!",
		"Hello, World.",
	}

	fmt.Println("Hello, World.")

	welcome(inputString[0], inputString[1])

	// debug

	fmt.Println(inputString[0])
}
