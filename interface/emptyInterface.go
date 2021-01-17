package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {
	v, ok := i.(int) //get the underlying int value from i
	fmt.Println(v, ok)
}

func main() {
	s := "Hello world!"
	i := 55
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	var x interface{} = "Steven Paul"
	assert(x)

	describe(s)
	describe(i)
	describe(strt)

}
