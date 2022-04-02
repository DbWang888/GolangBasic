package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写函数，接收文件路径及返回writer以及reader
func CopyFile(dstFileName string, srcFileName string) (writen int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("srcfile open err= ", err)
		return
	}
	defer srcFile.Close()
	//通过srcfile 获取 reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName 若不存在则创建
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("open dstfile err ", err)
		return
	}

	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)

}

func main() {
	//调用CopyFile完成

	srcFilePath := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\13_file\\07_fileWrite05\\test02.txt"
	dstFilePath := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\13_file\\08_fileCopy01\\tsetCopy.txt"

	a, err := CopyFile(dstFilePath, srcFilePath)

	if err == nil {
		fmt.Println("拷贝完成")
		fmt.Printf("%v", a)
	} else {
		fmt.Printf("%v", err)
	}

}
