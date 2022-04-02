package main

import (
	"fmt"
	"io/ioutil"
)

//不带缓冲区方式读取文件，适用于较小文件
//一次性读取文件

func main() {
	file := "C:/Users/wangli/Desktop/test/name.txt"
	content, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("read file err = %v", err)
	}

	//读取到的内容是[]byte，需要转成string输出
	fmt.Printf("%v", string(content))

}
