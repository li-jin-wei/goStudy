package main

import "fmt"

func main() {
	//	map结合silce使用02
	//通过map和slice 写一个简单的图书数据库，并可以通过书名或者书号来查询书的位置

	book1 := make(map[string]string)
	book1["num"] = "001"
	book1["name"] = "汉宣帝传"
	book1["addr"] = "A区第三排二书柜一层"

	book2 := make(map[string]string)
	book2["num"] = "002"
	book2["name"] = "汉朝风云"
	book2["addr"] = "C区第四排一书柜四层"

	book3 := make(map[string]string)
	book3["num"] = "003"
	book3["name"] = "中国历史"
	book3["addr"] = "A区第一排一书柜一层"

	book4 := map[string]string{"num": "004", "name": "go基础", "addr": "B区第四排四书柜一层"}

	books := make([]map[string]string, 0, 5)
	books = append(books, book1)
	books = append(books, book2)
	books = append(books, book3)
	books = append(books, book4)

	for _, book := range books {
		fmt.Println("图书的编号:", book["num"])
		fmt.Println("图书的地址", book["addr"])
	}

}
