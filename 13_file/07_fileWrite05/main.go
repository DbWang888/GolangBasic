package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//将test01 的文件内容写到test02中
//实际内容为清空后重新写入
//判断文件或者目录存在

//判断文件或者目录存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //代表文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func main() {
	file1Path := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\13_file\\07_fileWrite05\\test01.txt"
	file2Path := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\13_file\\07_fileWrite05\\test02.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("文件读取错误，", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Println("文件写入错误", err)
	}

	////判断文件或者目录存在
	ex, err := PathExists(file1Path)
	if ex == true {
		fmt.Println("file1存在")
	}

}
