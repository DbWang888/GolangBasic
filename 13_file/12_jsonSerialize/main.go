package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Person struct {
	Name  string `json:"name"` //tag标签
	Age   int
	Sal   float64
	Skill string
}

//序列化结构体为json格式
func testStruct() {

	person := Person{
		Name:  "老王",
		Age:   11,
		Sal:   66.66,
		Skill: "打麻将",
	}
	//将person序列化

	data, err := json.Marshal(&person)
	if err != nil {
		fmt.Println("序列化失败，", err)
		return
	}

	fmt.Println(string(data))
}

//将map进行序列化
func testMap() {
	//操作方法和结构体序列化一样

}

func main() {
	testStruct()
}
