package main

import "fmt"

func main() {

	//作为容器的一种，数组能够存储一组特定类型的数据，但是缺点是数组是定长的。系统根据定义的固定的长度，开辟了固定的内存大小所以不能改变大小
	//切片也是存储相同类型的树结构，但是不同于数组的是它的大小可以改变，如果长度不够也可以自动扩充

	//	声明切片
	var slice []int

	fmt.Println(slice)

	//	通常情况下使用map来创建切片
	nums := make([]int, 3, 5)
	fmt.Println(nums)
	fmt.Println(len(nums), cap(nums))

	//使用append函数追加元素
	//切片会自动扩容
	nums = append(nums, 4, 5, 6)
	fmt.Println(nums)
	fmt.Println(len(nums), cap(nums))

	num := make([]int, 0, 10)

	//表示将另一个切片数组完整加入到当前切片中
	num = append(num, nums...)
	fmt.Println(num)
}
