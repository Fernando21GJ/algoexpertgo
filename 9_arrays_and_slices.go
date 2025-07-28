package main

import "fmt"

func main() {
	/*	var arr [5]int
		arr[0] = 100
		fmt.Println(len(arr))

		arr2 := [2]int{4, 5}
		fmt.Println(len(arr2))

		arr3 := [...]int{4, 5, 7, 7, 6, 6, 2}
		fmt.Println(len(arr3))

		arr4 := [2][2]string{{"hello", "world"}, {"code", "go"}}
		fmt.Println(arr3, arr4)

		fmt.Printf("%T\n", arr4)

		for i, nestedArr := range arr4 {
			fmt.Printf("%d:%v\n", i, nestedArr)
		}

		for _, nestedArr := range arr4 {
			for j, nestedArr2 := range nestedArr {
				fmt.Printf("%d:%v\n", j, nestedArr2)
			}
		}

		test(arr4)
		fmt.Println(arr4)

		arr5 := [5]int{1, 2, 3, 4, 5}
		sliceArr := arr5[:4]
		fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))
		sliceArr = sliceArr[1:]
		fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))

		arr6 := [5]int{1, 2, 3, 4, 5}
		sliceArr = arr6[1:4]
		fmt.Println(arr6, sliceArr)
		sliceArr[0] = 100
		arr6[3] = 200
		fmt.Println(arr6, sliceArr)

		arr7 := []string{"hello", "world"}
		arr7 = arr7[:1]
		arr7 = arr7[:2]
		fmt.Println(arr7, len(arr7), cap(arr7))

		arr8 := []string{}

		for i := 0; i < 10; i++ {
			arr8 = append(arr8, "hello")
			fmt.Println(arr8, len(arr8), cap(arr8))
		}

		arr9 := make([][]int, 5, 10)
		fmt.Println(arr9, len(arr9), cap(arr9))

	*/

	arr2 := []bool{true, false, false}

	for i, val := range arr2 {
		fmt.Println(i, val)
	}
	arr3 := []bool{true, false, false}
	test(arr3)
	fmt.Println(arr3)
}

func test(arr []bool) {
	arr[0] = false
}

/*
func test(x [2][2]string) {
	x[0][0] = "new string"
}
*/
