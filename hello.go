package main // Defines the entry point. It tells Go that this file should compile as an
// executable program, not a shared library.

import "fmt" // Includes a library. Short for "format". It provides functions
// formatting text and printing it to the console.

// The starting function. This is where the program execution begins.
func main() {
	// Outputs text. It prints the string to the screen and adds a new line at the end.
	fmt.Println("Hello, World!")
}
