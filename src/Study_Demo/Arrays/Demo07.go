package main

import "fmt"

// 值传递
func main() {
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)

	//修改arr2的值，arr1的值不会改变
	//数组是值类型，拷贝一个新的内存空间
	arr2[0] = 10
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)

}
