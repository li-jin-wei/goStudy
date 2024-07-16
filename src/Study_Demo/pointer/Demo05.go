package main

import "fmt"

func main() {
	ptr := f5()

	fmt.Println("ptr:", ptr)
	fmt.Printf("ptr类型：%T\n", ptr)
	fmt.Println("ptr的地址:\n", &ptr)
	fmt.Println("ptr地址中的值:\n", *ptr)

	fmt.Println((*ptr)[0])
	ptr[0] = 10
	fmt.Println(ptr[0])
}

func f5() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	return &arr
}
