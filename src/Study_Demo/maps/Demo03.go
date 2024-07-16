package main

import "fmt"

func main() {
	//	map的遍历
	map1 := make(map[int]string)

	map1[1] = "张无忌"
	map1[2] = "张三丰"
	map1[3] = "常遇春"
	map1[4] = "胡青牛"

	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
