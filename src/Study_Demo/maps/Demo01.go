package main

import "fmt"

func main() {
	//map是go语言中的内置的字典类型，他存储的是一个键值对 key:value 类型的数据。
	//map也是一种容器，和数组切片不一样的是，他不是通过下标来访问数据，而是通过key来访问数据。

	//	map 特点
	//	map是无序的、长度不固定、不能通过下标获取，只能通过key来访问。
	//	map的长度不固定 ，也是一种引用类型。可以通过内置函数 len(map)来获取map长度。
	//	创建map的时候也是通过make函数创建。
	//	map的key不能重复，如果重复新增加的会覆盖原来的key的值。

	//声明map数据类型
	var m1 map[int]int
	fmt.Println(m1)

	//使用make声明
	m2 := make(map[string]int)
	fmt.Println(m2)

	//声明map数据类型并初始化赋值
	m3 := map[string]int{"语文": 90, "数学": 100, "英语": 40}
	fmt.Println(m3)

	//for k, v := range m3 {
	//	fmt.Println(k, v)
	//}
}
