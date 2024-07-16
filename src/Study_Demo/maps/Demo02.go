package main

import "fmt"

func main() {
	var m1 map[int]string
	var m2 = make(map[int]string)
	m3 := map[string]int{"语文": 89, "数学": 23, "英语": 90}

	fmt.Println(m1 == nil)
	fmt.Println(m2 == nil)
	fmt.Println(m3 == nil)

	if m1 == nil {
		m1 = make(map[int]string)
	}
	m1[1] = "小猫"
	m1[2] = "小猪"

	fmt.Println(m1[2])
	val, ok := m1[1]
	fmt.Println(val, ok) //结果返回两个值，一个是当前获取的key对应的val值。二是当前值是否存在，会返回一个false或者true

	//修改key对应的value值，如果值不存在，直接添加
	m1[1] = "小狗"
	fmt.Println(m1[1])

	//删除map值
	delete(m1, 1) //delete(map,key)
	fmt.Println(m1[1])
	fmt.Println(len(m1))
}
