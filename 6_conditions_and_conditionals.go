package main

import (
	"fmt"
)

func main() {
	a := 1.2
	b := 1

	b2 := int(a) == b
	c := b2
	fmt.Println(c)
}
