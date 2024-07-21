package main

import "fmt"

//切片的遍历

func main() {
	s1 := make([]int, 3, 5)
	fmt.Println(s1)

	s1 = append(s1, 1, 2, 3)

	//第一种方式，根据len()函数获取切片的长度，作为临界点
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	fmt.Println("-------------------------------------------------------------------")
	//第一种方式，根据range
	for j := range s1 {
		fmt.Println(s1[j])
	}
}
