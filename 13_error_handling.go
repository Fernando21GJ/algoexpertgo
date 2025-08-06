package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

/*
func doPanic() {
	//panic("fail")
}
*/
/*
func handlePanic() {
	r := recover()
	if r != nil {
		fmt.Println(r)
	}
	//fmt.Println("recover", r)
}*/
func main() {
	/*defer fmt.Println("defer")
	fmt.Println("Hello")
	panic("fail")
	fmt.Println("goodbye")*/
	/*defer handlePanic()
	fmt.Println("start")
	doPanic()
	fmt.Println("end")**/

	/*panic("test")
	r := recover()
	fmt.Println(r)*/

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result, err)
	}

}
