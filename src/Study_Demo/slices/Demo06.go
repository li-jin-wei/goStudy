package main

import "fmt"

func main() {
	//	切片的删除

	//获取切片指定位置的值，重新赋值给当前切片
	slice1 := []int{1, 2, 3, 4}
	//slice1 = slice1[1:]
	//fmt.Println(slice1)

	//使用append不会改变当前切片的内存地址
	//slice1 = append(slice1[:0], slice1[1:]...)
	//fmt.Println(slice1)

	//删除指定的下标元素
	i := 2 //要删除的下标为2
	slice1 = append(slice1[:i], slice1[i+1:]...)
	fmt.Println(slice1)

	//删除切片结尾的方法
	slice1 = slice1[:len(slice1)-1] //删除最后一个元素
	fmt.Println(slice1)
}
