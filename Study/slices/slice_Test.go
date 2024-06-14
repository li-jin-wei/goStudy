package main

import "fmt"

func main() {
	//var s1 []int
	//var s2 = []int{1, 2, 3}
	//fmt.Println(s1, s2)
	//fmt.Println(len(s1), cap(s1))
	//fmt.Println(len(s2), cap(s2))
	//s2 = append(s2, 20, 30, 1, 1, 1, 1)
	//fmt.Println(len(s2), cap(s2))

	var slice1 = make([]int, 0, 20)
	fmt.Println(len(slice1), cap(slice1))

	for i := 0; i < 30; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))

}
