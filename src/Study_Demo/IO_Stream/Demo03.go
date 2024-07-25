package main

import (
	"fmt"
	"os"
)

// IO读取文件
func main() {
	file, _ := os.Open("/Users/dayu/Desktop/test.txt")
	defer file.Close()

	//创建缓存池，一个汉字占3个字符
	//bs := make([]byte, 3, 1024)
	bs := make([]byte, 2, 1024)
	n, err := file.Read(bs)
	fmt.Println(n, err)
	fmt.Println(string(bs))

	file.Read(bs)
	fmt.Println(string(bs))

	file.Read(bs)
	fmt.Println(string(bs))

	file.Read(bs)
	fmt.Println(string(bs))
}
