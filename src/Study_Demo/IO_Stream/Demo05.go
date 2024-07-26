package main

import (
	"fmt"
	"io"
	"os"
)

/*
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
	offset 偏移量
	whence 如何设置，当前光标的位置

const (
	SeekStart   = 0  表示相对于文件的开头
	SeekCurrent = 1  表示相对于当前光标所在的位置
	SeekEnd     = 2  相对于文件的末尾
)

*/

func main() {
	file, _ := os.OpenFile("/Users/dayu/Desktop/test.txt", os.O_RDONLY, os.ModeAppend)
	defer file.Close()

	//构建缓存池
	buf := []byte{0}

	file.Seek(2, io.SeekStart)
	file.Read(buf)
	fmt.Println(string(buf))

	file.Seek(0, io.SeekCurrent)
	file.Read(buf)
	fmt.Println(string(buf))

	file.Seek(0, io.SeekEnd)
	file.WriteString("hello world")

	//file.Seek(5, io.SeekEnd)
	//file.Read(buf)
	//fmt.Println(string(buf))
}
