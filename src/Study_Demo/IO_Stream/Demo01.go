package main

import (
	"fmt"
	"os"
)

// 文件状态
func main() {
	fileinfo, err := os.Stat("/Users/dayu/Desktop/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	//Name() 文件名
	fmt.Println(fileinfo.Name())

	//Size()大小
	fmt.Println(fileinfo.Size())

	//ModTime() 修改时间
	fmt.Println(fileinfo.ModTime())

	//判断文件是否为目录，返回一个布尔值
	fmt.Println(fileinfo.IsDir())

	//Mode()文件权限
	fmt.Println(fileinfo.Mode())
}
