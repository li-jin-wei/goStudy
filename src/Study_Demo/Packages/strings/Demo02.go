package main

import (
	"fmt"
	"time"
)

//time包

func main() {
	timeDemo1()
	timeDemo2()
	timeDemo3()
}
func timeDemo1() {
	//获取当前时间now
	now := time.Now()
	//fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	//Printf ： 整数补位--02如果不足两位，左侧用0补齐输出
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timeDemo2() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

}

func timeDemo3() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	timeStr := "2024-07-18 16:51:30"
	timOBJ, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	fmt.Printf("%T\n", timOBJ)
	fmt.Println(timOBJ)
}
