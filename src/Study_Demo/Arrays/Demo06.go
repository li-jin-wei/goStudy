package main

import (
	"fmt"
)

// 多维数组的应用
func main() {
	var score [3][5]float64

	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			fmt.Printf("请输入%d班第%d个同学的成绩:\n", i+1, j+1)
			fmt.Scanln(&score[i][j])
		}
	}
	totalsum := 0.0
	for i := 0; i < len(score); i++ {
		sum := 0.0
		for j := 0; j < len(score[i]); j++ {
			sum += score[i][j]
		}
		totalsum += sum
		fmt.Printf("%d班的总分为:%v,平均分为:%v\n", i+1, sum, sum/float64(len(score[i])))
	}
	fmt.Printf("所有班级的总分为:%v，所有班的平均分为:%v\n", totalsum, totalsum/15)
	fmt.Println(score)
}
