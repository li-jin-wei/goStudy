package main

import "fmt"

func main() {
	//	数组的定义
	var num [5]int

	//	数组访问使用下标,并给数组赋值
	num[0] = 1
	num[1] = 10

	fmt.Println(num)

	//通过下标获取数组数值
	fmt.Println("数组第二个数值为:", num[1])

	//通过len函数获取数组长度
	fmt.Println("数组num的长度为:", len(num))

	//通过cap函数获取数组容量，因为数组是固定长度，所哟数组的长度和容量是相同的
	fmt.Println("数组num的容量为:", cap(num))

	//通过下标修改数组的值
	num[0] = 100
	fmt.Println(num[0])
}
