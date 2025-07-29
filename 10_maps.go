package main

import "fmt"

func main() {
	/*var mp map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	var mp2 map[string]int
	mp3 := map[string]int{"one": 1, "two": 2, "three": 3}

	mp4 := map[string][]int{"h": {1}, "a": {2}}
	mp5 := make(map[int]rune)
	mp5[5] = 2
	delete(mp5, 2)

	value, ok := mp5[6]

	mp6 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	fmt.Println(mp)
	fmt.Println(mp2)
	fmt.Println(mp3)
	fmt.Println(mp4)
	fmt.Println(mp5)
	fmt.Println(value, ok)

	for key, value := range mp6 {
		fmt.Println(key, value)
	}*/

	n := 100
	divisibleMap := make(map[int]uint)

	for num := 1; num <= n; num++ {
		for d := 1; d <= 5; d++ {
			if num%d == 0 {
				divisibleMap[d]++
			}
		}
	}
	fmt.Println(divisibleMap)
}
