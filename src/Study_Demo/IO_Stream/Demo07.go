package main

import (
	"fmt"
	"log"
	"os"
)

/*
遍历文件夹

	1.读取当前文件夹下的所有文件
	2.如果是文件夹，进入文件夹，继续读取里面的所有文件
	3.设置一些结构化代码
*/
func main() {
	dir := "/Users/dayu/Desktop/AD域"
	tree(dir, 0)
}

func tree(dir string, level int) {

	//编写层级
	//|--|--|--|--
	tabString := "|--"
	for i := 0; i < level; i++ {
		tabString = tabString + "|--"
	}

	//获取目录ReadDir，返回目录信息[]DirEntry,多个文件信息
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}

	//遍历出所有文件之后，获取里面的单个文件
	for _, file := range fileInfos {
		//文件夹中文件的全路径展示
		fileName := dir + "/" + file.Name()
		fmt.Println(tabString + file.Name())
		//如果是文件夹，再次遍历，才用递归函数
		if file.IsDir() {
			tree(fileName, level+1)
		}
	}
}
