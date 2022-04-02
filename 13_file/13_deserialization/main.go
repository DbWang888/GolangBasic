package main

import (
	"encoding/json"
	"fmt"
)

//json字符串反序列化为map、struct、slice等
//反序列化之前需要先声明对应的map、struct、slice等

type Person struct {
	Name  string
	Age   int
	Sal   float64
	Skill string
}

//反序列化函数，将json格式字符串转为byte切片，再进行反序列化
func UmarshalStruct() {
	str := `{"name":"老王","Age":11,"Sal":66.66,"Skill":"打麻将"}`

	//定义一个Person实例

	var person Person
	var personMap map[string]interface{} //map不需要再make  Unmarshal函数底层集成了make操作

	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Printf("err = %v", err)
	}
	fmt.Printf("反序列化后%+v", person)

	err = json.Unmarshal([]byte(str), &personMap)
	if err != nil {
		fmt.Printf("err = %v", err)
	}
	fmt.Printf("反序列化后%+v", personMap)

}

func main() {
	UmarshalStruct()
}
