package main

import "time"

func main() {

	//实例化Operate对象
	operate := Operate{
		DirectoryPath:          "/Users/dayu/Desktop/Demo_Test",
		RemoveAllDirectoryPath: "/Users/dayu/Desktop/Demo",
		CreateAllDirectoryPath: "/Users/dayu/Desktop/Demo/1/2/3/4",
		CreateFilePath:         "/Users/dayu/Desktop/test.txt",
		DeleteFilePath:         "/Users/dayu/Desktop/test.txt",
	}

	operate.RemoveDirectory(operate.DirectoryPath)
	operate.RemoveAllDirectory(operate.RemoveAllDirectoryPath)
	//休息三秒查看执行结果
	time.Sleep(5 * time.Second)

	operate.MakeDirectory(operate.DirectoryPath)
	operate.MakeAllDirectory(operate.CreateAllDirectoryPath)

	operate.FileCreate(operate.CreateFilePath)

	operate.FileDelete(operate.DeleteFilePath)

}
