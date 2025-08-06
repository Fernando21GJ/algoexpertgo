package main

import "fmt"

func test1(a int, b int) (x int, y int) {
	x = a
	y = b
	return
}

func sumInts(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	//fmt.Println(nums)
	return sum
}

func returnFunc(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func callFunc(callable func(string) string, arg string) {
	result := callable(arg)
	fmt.Println(result)
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	a, b := test1(10, 20)
	fmt.Println(a, b)

	//sum := sumInts(1, 2, 3, 4, 5, 6)
	sum := sumInts(nums...)
	fmt.Println(sum)

	fn := returnFunc(100)
	value := fn(50)
	fmt.Println(value)

	pos := adder()
	neg := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	myFunc := func(str string) string {
		return str + "hello!"
	}
	callFunc(myFunc, "world")
}
