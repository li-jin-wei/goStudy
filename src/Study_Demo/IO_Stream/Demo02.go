package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	MkdirDemo()
	time.Sleep(3 * time.Second)
	MkdirAllDemo()
}

func MkdirDemo() {
	//Mkdir()创建目录
	err := os.Mkdir("/Users/dayu/Desktop/Demo_Test", os.ModePerm)
	if err != nil {
		fmt.Println(err, "文件存在")
	}
	fmt.Println("文件创建完毕")
}

func MkdirAllDemo() {
	//MkdirAll()层级文件创建
	err := os.MkdirAll("/Users/dayu/Desktop/Demo/1/2/3/4", os.ModePerm)
	if err != nil {
		fmt.Println(err, "文件存在")
	}
	fmt.Println("层级文件创建完毕")
}

func RemoveDemo() {

}
