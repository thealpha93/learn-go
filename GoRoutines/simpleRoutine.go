package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello Everyone")
}

// Main function will run on the `main` go routine
func main() {
	// Create a new go routine
	// GO control doesn't wait for the the routine to finish
	// The control is passed directly to the next line
	// Any return values from a routine is ignored
	// If the main go routine terminates, no other go routine will run
	go hello()
	time.Sleep(1 * time.Second)

	// Only this is printed if no sleep is provided
	fmt.Println("Main function")
}
