package main

import (
	"fmt"
	"strings"
)

//strings 字符串常用操作包

func main() {
	str := "Hello_World"
	str2 := "Test"
	//Contains 判断某个字符是否包含了指定的内容
	fmt.Println(strings.Contains(str, "world"))

	//ContainsAny 判断参数一是否有参数二的任意字符
	fmt.Println(strings.ContainsAny(str, str2))

	//Count 统计字符串中某个字符出现的次数
	fmt.Println(strings.Count(str, "l"))

	fileName := "20230219.mp3"
	//HashPrefix 判断用什么开头
	if strings.HasPrefix(fileName, "2023") {
		fmt.Println("找到2023开头的文件:", fileName)
	}
	//HaSuffix 判断用什么结尾
	if strings.HasSuffix(fileName, ".mp4") {
		fmt.Println("找到mp4结尾的文件:", fileName)
	}

	//Index() 寻找这个字符串第一次出现的位置
	fmt.Println(strings.Index(str, "d"))

	//LastIndex() 寻找这个字符串最后一次出现的位置 LastIndex()
	fmt.Println(strings.LastIndex(str, "o"))

	//Join() 拼接字符串，数组或者切片拼接
	str3 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(strings.Join(str3, ""))

	//Split()，通过某个格式拆分字符串,返回的是一个字符串切片
	str4 := strings.Join(str3, "|")
	fmt.Println(strings.Split(str4, "|"))

	//ToUpper()大写
	fmt.Println(strings.ToUpper(str))
	//ToLower()小写
	fmt.Println(strings.ToLower(str))

	//Replace 替换
	fmt.Println(strings.Replace(str, "H", "new", 1))

	//n代表替换几次
	fmt.Println(strings.Replace(str, "l", "t", 1))

}
