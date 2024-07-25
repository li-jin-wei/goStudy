package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, err := os.Stat("/Users/dayu/Desktop/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.ModTime())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Size())
}
