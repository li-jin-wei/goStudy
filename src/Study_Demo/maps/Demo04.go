package main

import "fmt"

func main() {
	//	map结合silce使用01

	user1 := make(map[string]string)
	user1["name"] = "lisi"
	user1["age"] = "18"
	user1["sex"] = "男"
	user1["addr"] = "蜀山"

	user2 := make(map[string]string)
	user2["name"] = "wanger"
	user2["age"] = "20"
	user2["sex"] = "女"
	user2["addr"] = "成都"

	user3 := map[string]string{"name": "小蓝", "age": "18", "sex": "男", "addr": "火星"}

	userDatas := make([]map[string]string, 0, 3)
	userDatas = append(userDatas, user1)
	userDatas = append(userDatas, user2)
	userDatas = append(userDatas, user3)

	//fmt.Println(userDatas)

	for _, user := range userDatas {

		fmt.Println("姓名:", user["name"])
		fmt.Println("年龄:", user["age"])
		fmt.Println("地址:", user["addr"])
		fmt.Println()
	}
}
