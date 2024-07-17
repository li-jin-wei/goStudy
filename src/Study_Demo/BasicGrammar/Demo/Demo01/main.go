package main

import (
	"fmt"
	"strconv"
)

var s5 int64 = 128

func main() {
	var s1 = 77
	var s2 = "99"
	var s3 = true
	var s4 = -111

	fmt.Println("隐式定义变量")
	fmt.Printf("%T,%T,%T,%T\n", s1, s2, s3, s4)
	fmt.Println("---------------------类型转换-------------------------")

	i1, _ := strconv.Atoi(s2)
	i2, _ := strconv.ParseInt(s2, 10, 64)
	i3 := strconv.Itoa(s1)
	i4 := strconv.FormatInt(s5, 36)
	i5, err := strconv.ParseFloat(s2, 128)

	i6, _ := strconv.ParseFloat(s2, 64)
	fmt.Println(i1, i2, i3, i4, i5)
	fmt.Printf("%T,%T,%T,%T,%T,%T\n", i1, i2, i3, i4, i5, i6)
	fmt.Println(err)

}
