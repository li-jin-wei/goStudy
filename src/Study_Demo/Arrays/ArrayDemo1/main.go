package main

import "fmt"

func main() {
	//nums := [3]int{1, 2, 3}
	//num := new([10]int)
	//a := [...]int{9, 8, 7, 6}
	//
	//num[2] = 59
	//
	//for i := 0; i < len(nums); i++ {
	//	fmt.Println(nums[i])
	//}
	//
	////i是数组下标，
	//for i, v := range a {
	//	fmt.Println(i, v)
	//}
	//fmt.Println(nums, num[2], len(a), cap(a))
	//	len是切片长度，cap是切片容量

	// 多维数组
	p := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	//for i := 0; i < len(p); i++ {
	//	for j := 0; j < len(p[i]); j++ {
	//		fmt.Printf("%v\n", p[i][j])
	//		//	p[0][0],p[0][1],p[0][2]
	//		//	p[1][0],p[1][1],p[1][2]
	//	}
	//}

	//在p这个多维数组中，取索引和值，v是第一层数组，即 【1，2，3】，【4，5，6】
	for i, v := range p {
		for j, v2 := range v {
			fmt.Println(i, j, v2)
		}
	}
}
