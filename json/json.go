package main

import (
	"encoding/json"
	"fmt"
)

type Json struct {
	Id      int
	Order   string   `json:"order"` //二次编码
	Subject []string // `json:"-"`  不输出到屏幕
	Isok    bool     `json:",string"`
	Price   float64
}

func main() {
	// 结构体生成json
	s := Json{1, "GW123123", []string{"Go", "php"}, true, 453.2}
	// 编码 根据内容生成json
	// buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(string(buf))

	// map生成json   创建map
	m := make(map[string]interface{}, 4)
	m["id"] = 1
	m["order"] = "GW43524"
	m["subject"] = []string{"go", "php"}
	m["isok"] = true
	m["price"] = 347.65
	// 生成json
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println("map json:", string(res))

	// json解析到结构体
	jsonStr := `
{
	"Id": 1,
	"order": "GW123123",
	"Subject": [
		"Go",
		"php"
	],
	"Isok": "true",
	"Price": 453.2
}`

	// 定义一个结构体
	var tmp Json
	errs := json.Unmarshal([]byte(jsonStr), &tmp)
	if errs != nil {
		fmt.Println("err", errs)
		return
	}
	fmt.Println("json解析到结构体:", tmp)
}
