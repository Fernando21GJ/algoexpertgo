package main

import "fmt"

func main() {
	/*var a int = 20
	switch a {
	case 10:
		fmt.Println("a is 10")
	default:
		fmt.Println("a is  not 10")
	}*/

	/*a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")

	}*/

	/*a := -1 //3
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}*/

	/*switch {
	case false: //1 < 2
		fmt.Println("one")
	case 2 < 3:
		fmt.Println("two")
	default:
		fmt.Println("default")

	}*/
	/*switch a := -2; {
	case a < -1:
		fmt.Println("a less than -1")
		fallthrough
	case a < 0:
		fmt.Println("a less than 0")
		fallthrough
	case a < 1:
		fmt.Println("a less than 1")
		default:
		fmt.Println("default")
	}*/
	/*a := "h"
	switch a {
	case "h", "a", "j":
		fmt.Println("worked")
	}*/
	switch a := "h"; {
	case a == "h", a == "e":
		fmt.Println("worked")
	}
}
