package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a := 5
	b := int8(2)
	c := float64(a) / float64(b)

	d := 5
	d = d + 10
	e := math.Max(5, 10)
	f := math.Sqrt(21)
	g := math.Pow(10, 2)
	h := math.Round(20.87)
	i := math.Ceil(20.7)
	j := math.Floor(20.8)

	str := "1234"
	str2 := "12345Hello"
	str3 := "123456"

	k, err := strconv.Atoi(str)
	l, err := strconv.Atoi(str2)
	m, err := strconv.ParseInt(str3, 10, 64)

	n := int8(10)
	o := int64(1272727)
	p := int64(n) + o

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println(k, err)
	fmt.Println(l, err)
	fmt.Println(m, err)
	fmt.Println(p)
}
