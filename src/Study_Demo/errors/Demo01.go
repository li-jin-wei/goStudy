package main

import (
	"fmt"
	"os"
)

// 错误是开发中必须要思考的问题
// - 某些系统错误 ，文件被占用，网络有延迟
// - 人为错误：核心就是一些不正常的用户会怎么来给你传递参数，sql注入

func main() {
	file, err := os.Open("test.txt")
	// 在开发中，我们需要思考这个错误的类型  PathError
	// 1、文件不存在 err
	// 2、文件被占用 err
	// 3、文件被损耗 err
	// 调用方法后，出现错误，需要解决

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name())
}
