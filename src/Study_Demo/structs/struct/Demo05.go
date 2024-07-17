package main

import "fmt"

type User2 struct {
	name string
	age  int
	sex  string
}

func main() {
	//	结构体指针
	user1 := User2{"dd", 18, "男"}
	fmt.Println(user1)
	fmt.Printf("%T,%p\n", user1, &user1)

	user2 := user1
	fmt.Println(user2)
	fmt.Printf("%T,%p\n", user2, &user2)

	//结构体是值类型的
	user2.name = "tywin"
	fmt.Println(user1)
	fmt.Println(user2)

	var userPtr *User2
	userPtr = &user1
	fmt.Println(*userPtr)
	(*userPtr).name = "DDDD"
	fmt.Println(user1)

	fmt.Println("---------------------")
	userPtr.name = "张三"
	fmt.Println(user1)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++")
	//内置函数 new 创建对象。  new 关键字创建的对象，都返回指针，而不是结构体对象。
	user3 := new(User2)
	fmt.Println(user3)
	fmt.Println(*user3)

	(*user3).name = "小红"
	user3.age = 18
	user3.sex = "女"
	fmt.Println(*user3)
}
