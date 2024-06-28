package main

import (
	"fmt"
	"os"
	"time"
)

type FileInfo interface {
	Name() string
	Size() int64
	Mode() os.FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() any
}

func main() {
	fileinfo, err := os.Lstat("/Users/dayu/Desktop/备忘.txt")
	if err != nil {
		return
	}
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.ModTime())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Sys())
}
