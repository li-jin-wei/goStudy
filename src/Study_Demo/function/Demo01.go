package main

import "fmt"

// copy
func main() {
	//是将funcA赋值给了h,
	//并不是接受了funcA的返回值，所以结果不为nil
	h := funcA
	if h == nil {
		fmt.Println("h is nil")
	} else {
		fmt.Println("h is not nil")
	}
}

func funcA() []string {
	return nil
}
