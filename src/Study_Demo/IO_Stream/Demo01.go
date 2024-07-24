package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, err := os.Lstat("/Users/dayu/Desktop/备忘.txt")
	if err != nil {
		return
	}

	//文件名称
	fmt.Println(fileinfo.Name())
	//文件大小
	fmt.Println(fileinfo.Size())
	//文件权限
	fmt.Println(fileinfo.Mode())
	//文件修改模式
	fmt.Println(fileinfo.ModTime())
	//是否为目录
	fmt.Println(fileinfo.IsDir())

}
