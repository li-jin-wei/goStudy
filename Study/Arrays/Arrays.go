package main

import "fmt"

//多维数组

//var arrays  [2][4]int

var arrays = [2][4]int{
	[4]int{1, 2, 3, 4},
	[4]int{9, 8, 7, 6},
}

func main() {
	fmt.Println(arrays)
	for _, v1 := range arrays {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
