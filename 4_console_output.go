package main

import "fmt"

func main() {
	//fmt.Println("XDDDD")
	//var x int = 1
	//var y bool = true
	//a := 3.4567
	//text := "ProgrammingExpert"

	i := 1
	b := 4.359
	c := i
	//z := 10.2716172
	//fmt.Printf("%vhello%T", x, y)
	//fmt.Printf("%T:%T", x, y)
	//fmt.Printf("%.2f", a)
	//fmt.Printf("\"%s\"", text)

	z := fmt.Sprintf("%.2f + %d", b, i)
	fmt.Println(z, c)
	//str := fmt.Sprintf("%T", 1)
	//fmt.Println(str)
	//fmt.Printf("%v:%v", 1, true)
}

func DefineConstant(x string) (string, int) {
	return x, 1
}
func Format(a float64, b bool, c string) string {
	// Write your code here
	result := fmt.Sprintf("%%%v:%v/%s", a, b, c)
	return result
}
