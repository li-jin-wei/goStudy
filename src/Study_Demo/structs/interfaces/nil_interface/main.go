package main

import "fmt"

// 字典结构
type Dictionary struct {
	data map[string]interface{}
}

func (d *Dictionary) GetData(key string) interface{} {
	return d.data[key]
}
func (d *Dictionary) SetData(key string, value interface{}) {
	d.data[key] = value
}

func NewDict() *Dictionary {
	return &Dictionary{
		data: make(map[string]interface{}),
	}
}

func main() {
	dict := NewDict()
	dict.SetData("1111", 1234)
	dict.SetData("2222", "lllll")
	d := dict.GetData("1111")
	d1 := dict.GetData("2222")
	fmt.Println(d, d1)
}
