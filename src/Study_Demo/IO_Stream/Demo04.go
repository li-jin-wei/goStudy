package main

import (
	"fmt"
	"os"
)

//文件IO写入

func main() {
	file, _ := os.OpenFile("/Users/dayu/Desktop/test.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	bs := []byte{65, 66, 67, 68, 69}
	n, err := file.Write(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	n, err = file.WriteString("Test string Write")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
