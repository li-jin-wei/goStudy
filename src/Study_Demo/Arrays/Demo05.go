package main

import "fmt"

//多维数组的遍历

func main() {

	// 多维数组
	p := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			fmt.Printf("%v\n", p[i][j])
			//	p[0][0],p[0][1],p[0][2]
			//	p[1][0],p[1][1],p[1][2]
		}
	}

	//在p这个多维数组中，取索引和值，v是第一层数组，即 【1，2，3】，【4，5，6】
	for i, v := range p {
		for j, v2 := range v {
			fmt.Println(i, j, v2)
		}
	}
}
