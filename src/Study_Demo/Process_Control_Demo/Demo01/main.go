package main

import "fmt"

func main() {
	var a int = 60 // 00111100
	var b int = 13 // 00001101
	//按位与运算，都是1则为1，否则为0
	fmt.Println(a & b)
	//按位或运算，都是0则为0
	fmt.Println(a | b)
	//异或运算，不同为1，相同位0
	fmt.Println(a ^ b)
	//位清空，a&^b，对于b上的每个数值，如果为0，则取a对应位上的数值，如果为1，则取0
	fmt.Println(a &^ b)

	//左移 2位，最高位移动两位，低位用0补位
	//60 111100
	//左移两位  11110000 240
	fmt.Println(a << 2)
	//111100
	//001111 15
	fmt.Println(a >> 2)

	//右移2位，低位移动，高位用0补齐
	//13 001101
	//110100 32+16+4=52
	fmt.Println(b << 2)
	//右移两位 000011 3
	fmt.Println(b >> 2)

}
