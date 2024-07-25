package main

import (
	"log"
	"os"
)

func main() {
	//file, err := os.Open("/Users/dayu/Desktop/test.txt")
	//if err != nil {
	//	log.Println(err)
	//}
	//defer file.Close()

	//OpenFile 返回的是*file和 err
	file, err := os.OpenFile("/Users/dayu/Desktop/test.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	log.Println(file)
	defer file.Close()
}
