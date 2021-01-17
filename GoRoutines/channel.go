package main

import (
	"fmt"
)

func main() {
	var a chan int

	if a == nil {
		fmt.Println("Channel doesn't exist. Going to define it")
		a = make(chan int)

		fmt.Printf("Type of channel is %T", a)
	}
}
