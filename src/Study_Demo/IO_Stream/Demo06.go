package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

/*
	  断点续传
	1.需要记录上一次传递了多少数据、temp.txt 记录
	2.如果被暂停或者被中断了，就可以读取这个temp.txt的记录，恢复上传
	3.删除temp.txt
*/

func main() {
	//传输的源文件地址
	srcFile := "/Users/dayu/Desktop/client/sql_backup.sql"
	//传输的目标文件地址
	destFile := "/Users/dayu/server/upload_sql_backup.sql"
	//临时记录文件
	tempFile := "/Users/dayu/server/temp.txt"

	//创建对应的file对象，连接起来
	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file1.Close()
	defer file2.Close()
	fmt.Println("file1/2/3文件连接建立完毕")

	//读取temp.txt
	file3.Seek(0, io.SeekStart)
	buf := make([]byte, 1024, 1024)
	n, _ := file3.Read(buf)
	//转换成string - 数字
	countStr := string(buf[:n])
	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println("temp.txt中记录的值为:", count)

	//设置读写偏移量
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	fmt.Println("file1/2 光标已经移动到了目标位置")

	//开始读写(复制、上传)
	bufData := make([]byte, 1024, 1024)

	//需要记录读取了多少个字节
	total := int(count)

	for {
		//读取数据
		readNum, err := file1.Read(bufData)
		if err == io.EOF {
			fmt.Println("文件传输完成")
			file3.Close()
			os.Remove(tempFile)
			break
		}

		//向目标文件中写入数据
		writeNum, err := file2.Write(bufData[:readNum])

		//将写入数据放到total中，在这里total就是传输的进度
		total += writeNum
		file3.Seek(0, io.SeekStart) //将光标重制到开头
		file3.WriteString(strconv.Itoa(total))
	}
}
