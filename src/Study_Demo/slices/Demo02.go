package main

import "fmt"

func main() {
	//	make()是Go语言中的内置函数，主要用于创建并初始化slice切片类型，或者map字典类型，或者channel通道类型数据。他与new方法的区别是。
	//	new用于各种数据类型的内存分配，在Go语言中认为他返回的是一个指针。指向的是一个某种类型的零值。make 返回的是一个有着初始值的非零值

	slice1 := new([]int)
	fmt.Println(slice1) //&[]

	slice2 := make([]int, 10)
	fmt.Println(slice2) //[0 0 0 0 0 0 0 0 0 0]

	//fmt.Println(slice1[0])
	//invalid operation: cannot index slice1 (variable of type *[]int)

	fmt.Println(slice2[0]) //0
}
