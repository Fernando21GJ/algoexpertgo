package main

import "fmt"

/*
x := 'a'         // implicit
var y byte = 'a' // explicit
*/
func main() {
	y := "hello"

	fmt.Printf("%T", y)
}
