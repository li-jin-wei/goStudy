package main

import "fmt"

func main() {
	//指针数组

	a := 1
	b := 2
	c := 3
	d := 4

	arr1 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr1)
	fmt.Println(*arr1[0])

	*arr1[0] = 100
	fmt.Println(arr1)
	fmt.Println(*arr1[0])

	a = 200
	fmt.Println(*arr1[0])
	fmt.Println(arr1)

	//	疑问  为什么修改元素之后内存地址还是一样的
}
