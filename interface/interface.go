package main

import (
	"fmt"
)

// Interface definition
type vowelsFinder interface {
	FindVowels() []rune
}

// Create a new local type
// Cannot define new methods on non-local types
type myString string

// type myString implements vowelsFinder
func (ms myString) FindVowels() []rune {
	// Declare variable with it's type as rune
	var vowels []rune

	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := myString("Sam Anderson")
	var v vowelsFinder

	v = name // Posiible since myString implementss vowelsFinder
	fmt.Printf("vowels are %c", v.FindVowels())
}
