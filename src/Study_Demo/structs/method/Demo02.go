package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() {
	fmt.Printf("User name: %s, age: %d\n", u.Name, u.Age)
}

func main() {
	u1 := User{"张三", 20}
	u1.String()
}
