package main

import (
	"fmt"
	"os"
)

func main() {
	file1, err := os.Open("/Users/dayu/Desktop/备忘.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file1)
	defer func(file1 *os.File) {
		err := file1.Close()
		if err != nil {

		}
	}(file1)

}
