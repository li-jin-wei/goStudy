package main

import "fmt"

type User struct {
	name string
	age  int
	sex  string
}

func main() {
	//	实例化对象的方法合集

	//第一种实例化方式
	var user1 User
	user1.name = "张三"
	user1.age = 18
	user1.sex = "女"
	fmt.Println("user1", user1)
	fmt.Println("user1.name:", user1.name)

	//第二种实例化方式
	user2 := User{}
	user2.name = "李四"
	user2.age = 18
	user2.sex = "男"
	fmt.Println("user2", user2)
	fmt.Println("user2.age:", user2.age)

	//第三种实例化方式
	user3 := User{
		name: "wanger",
		age:  20,
		sex:  "nv",
	}
	fmt.Println("user3", user3)
	fmt.Println("user3.age:", user3.age)

	//第四种实例化方式，需要与参数一一匹配
	user4 := User{"guoguo", 100, "女"}
	fmt.Println("user4", user4)
}
