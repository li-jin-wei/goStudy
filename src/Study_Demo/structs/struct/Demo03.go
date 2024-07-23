package main

import (
	"encoding/json"
	"fmt"
)

type Prescripts struct {
	Name     string      `json:"name,omitempty"`
	Unit     string      `json:"unit"`
	Additive *Prescripts `json:"additive,omitempty"`
}

func main() {
	p := &Prescripts{}
	p.Name = "测试"
	p.Unit = "1.2"
	p.Additive = &Prescripts{
		Name: "Test",
		Unit: "0.5",
	}
	//结构体转json字符串
	marshal, err := json.Marshal(p)
	if err != nil {
		return
	}
	fmt.Println("json=", string(marshal))
}
