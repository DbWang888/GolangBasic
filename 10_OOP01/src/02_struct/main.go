package main
import (
	"fmt"
	"encoding/json"
) 

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}


func main(){
	ps1 := Person {"jack", 18, "吃饭"}
	jsonps1, err := json.Marshal(ps1)
	if err != nil {
		fmt.Println("json错误", err)
	}
	fmt.Println(string(jsonps1))
}