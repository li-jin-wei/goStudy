package main

import (
	"log"
	"os"
)

//定义结构体实现接口方法

type Operate struct {
	//单个文件路径
	DirectoryPath string
	//层级文件删除第一层文件路径
	RemoveAllDirectoryPath string
	//创建层级文件详细目录
	CreateAllDirectoryPath string

	//创建文件路径
	CreateFilePath string

	//删除文件路径
	DeleteFilePath string
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

func (operste Operate) MakeAllDirectory(CreateAllDirectoryPath string) {
	//MkdirAll()层级文件创建
	err := os.MkdirAll(CreateAllDirectoryPath, os.ModePerm)
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

func (Operate Operate) FileCreate(CreateFilePath string) {
	file, err := os.Create(CreateFilePath)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("文件创建成功", file.Name())
	}
}

func (Operate Operate) FileDelete(DeleteFilePath string) {
	err := os.Remove(DeleteFilePath)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("文件删除成功")
	}

}
