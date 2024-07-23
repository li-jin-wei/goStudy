package main

import (
	"fmt"
	"strconv"
)

// strvonv包的使用
func main() {
	StringDemo()
	IntDemo()
}

func StringDemo() {
	s1 := "true"

	//ParseBool()字符串转为布尔值
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//%t 布尔值占位，%T 输出数据类型
	fmt.Printf("%T, %t\n", s1, b1)

	s2 := strconv.FormatBool(b1)
	//%s 字符串占位
	fmt.Printf("%T,%s\n", s2, s2)
}

func IntDemo() {
	s := "10000"

	//ParseInt()字符串转int
	//str,进制，大小
	i, _ := strconv.ParseInt(s, 10, 64)
	//%d 十进制
	fmt.Printf("%T,%d\n", i, i)

	//FormatInt()int转字符串
	s1 := strconv.FormatInt(i, 10)
	fmt.Printf("%T,%s\n", s1, s1)

	//Atoi() 字符串转10进制
	atio, _ := strconv.Atoi("-20")
	fmt.Printf("%T,%d\n", atio, atio)

	//Itoa() 10进制转字符串
	itoa := strconv.Itoa(30)
	fmt.Printf("%T,%s\n", itoa, itoa)
}
