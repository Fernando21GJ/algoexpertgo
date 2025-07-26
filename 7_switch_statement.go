package main

import (
	"fmt"
)

func main() {
	var a int = 20
	switch a {
	case 10:
		fmt.Println("a is 10")
	default:
		fmt.Println("a is  not 10")
	}
}
