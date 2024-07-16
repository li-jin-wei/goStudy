package main

import "fmt"

func main() {
	//	初始化数组的几种方式

	// 在定义数组的时候就直接初始化
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	//快速初始化
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	//使用...设置不定长度的数组，根据存入的数据来决定数组长度和容量
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("数组arr3的长度:", len(arr3))
	fmt.Println("数组arr3的容量:", cap(arr3))

	//数组默认值，只想给其中的某几个index位置赋值
	var arr4 [10]int
	arr4 = [10]int{1: 100, 10, 1000}
	fmt.Println(arr4)

}
