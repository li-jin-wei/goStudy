package main

import "fmt"

func main() {
	//	数组的遍历

	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println("数组的第", i+1, "个值为:", arr[i])
	}

	for i, v := range arr {
		fmt.Println("数组的第", i+1, "个值为:", v)
	}
}
