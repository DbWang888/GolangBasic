package testdemo02

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m Monster) Store() {

	//将结构体变量序列化为json格式,得到byte切片
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败", err)
	}

	//保存入文件中
	filePath := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\14_unit_testing\\02_testdemo02\\montser.txt"
	file, errw := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	if errw != nil {
		fmt.Println("文件创建错误")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))

	writer.Flush()

}

func (m *Monster) ReStore() {
	filePath := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\14_unit_testing\\02_testdemo02\\montser.txt"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}

	// file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	// if err != nil {
	// 	fmt.Println("读取失败", err)
	// }
	// defer file.Close()

	// reader := bufio.NewReader(file)

	//读取字符串
	// str1 := ""
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(str)
	// 	str1 += str
	// }

	//反序列化
	err = json.Unmarshal(data, m)

	if err != nil {
		fmt.Println("反序列化错误", err)
	}

}
