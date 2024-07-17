package main

import (
	"fmt"
	"strings"
)

//strings 字符串常用操作包

func main() {
	str := "hello world"
	str2 := "Test"
	//Contains 判断某个字符是否包含了指定的内容
	fmt.Println(strings.Contains(str, "world"))

	//ContainsAny 判断参数一是否有参数二的任意字符
	fmt.Println(strings.ContainsAny(str, str2))

	//Count 统计字符串中某个字符出现的次数
	fmt.Println(strings.Count(str, "l"))

}
