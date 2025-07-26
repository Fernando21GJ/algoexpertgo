package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for j := 0; j > -10; j-- {
		fmt.Println(j)
	}
	for k := 0; k <= 10; k += 2 {
		fmt.Println(k)
	}

	l := 0
	for l <= 10 {
		l++
		fmt.Println("print l ", l)
	}
	str := "hello"
	//fmt.Println(str[0])
	//fmt.Println(string(str[0]))
	//fmt.Printf("%T", str[0])

	for m := 0; m < len(str); m++ {
		//fmt.Println(str[m])
		//fmt.Printf("%T", str[m])
		fmt.Printf("%c\n", str[m])
	}

	for n, char := range str {
		//fmt.Println(n, char)
		fmt.Println(n, string(char))
	}

	for _, char := range str {
		fmt.Printf("%T", char)
		break
	}
}
