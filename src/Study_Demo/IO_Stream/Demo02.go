package main

import (
	"log"
	"os"
	"time"
)

type DirectoryOperate interface {

	// MakeDirectory 创建单个目录
	MakeDirectory(path string)

	//MakeAllDirectory 创建层级目录
	MakeAllDirectory(path string)

	//RemoveDirectory 删除单个目录
	RemoveDirectory(path string)

	//RemoveAllDirectory 删除层级目录
	RemoveAllDirectory(path string)
}
type Operate struct {
	//单个文件路径
	DirectoryPath string
	//层级文件删除第一层文件路径
	RemoveAllDirectoryPath string
	//创建层级文件详细目录
	MakeAllDirectoryPath string
}

func (operste Operate) MakeDirectory(DirectoryPath string) {
	//Mkdir()创建目录
	err := os.Mkdir(DirectoryPath, os.ModePerm)
	if err != nil {
		log.Println("创建失败，目录已存在")
		return
	}
	log.Println("文件创建完毕")
}

func (operste Operate) MakeAllDirectory(MakeAllDirectoryPath string) {
	//MkdirAll()层级文件创建
	err := os.MkdirAll(MakeAllDirectoryPath, os.ModePerm)
	if err != nil {
		log.Println("创建失败，目录已存在")
		return
	} else {
		log.Println("层级目录创建成功")
	}
}

func (operste Operate) RemoveDirectory(DirectoryPath string) {
	//Remove 只能删除单个空文件夹

	err := os.Remove(DirectoryPath)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("文件删除成功")

}

func (Operate Operate) RemoveAllDirectory(RemoveAllDirectoryPath string) {
	//强制删除此路径下的所有东西，危险操作
	err := os.RemoveAll(RemoveAllDirectoryPath)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("层级文件删除成功")
}

func main() {
	operate := Operate{DirectoryPath: "/Users/dayu/Desktop/Demo_Test", RemoveAllDirectoryPath: "/Users/dayu/Desktop/Demo", MakeAllDirectoryPath: "/Users/dayu/Desktop/Demo/1/2/3/4"}
	//实例化

	operate.RemoveDirectory(operate.DirectoryPath)
	operate.RemoveAllDirectory(operate.RemoveAllDirectoryPath)

	//休息三秒查看执行结果
	time.Sleep(3 * time.Second)
	operate.MakeDirectory(operate.DirectoryPath)
	operate.MakeAllDirectory(operate.MakeAllDirectoryPath)

}
