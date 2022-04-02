package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//打开文件并读取显示后再写入

func main() {
	filePath := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\13_file\\06_fileWrite04\\test.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开错误，", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完成 ")
			break
		}
		fmt.Println(str)
	}

	str := "新增\n"
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}

	write.Flush()

}
