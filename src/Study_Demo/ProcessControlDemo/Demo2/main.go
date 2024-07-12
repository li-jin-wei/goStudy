package main

import "fmt"

// fallthrough 穿透switch 当前case执行完了 继续执行下一case 不用匹配下一case是否符合直接执行
// **fallthrough必须放在case最后一行
func main() {
	m := 2
	switch m {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	}
}
