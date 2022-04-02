package main

import (
	"bufio"
	"fmt"
	"os"
)

//打开一个已经存在的文件，并在文件末尾追加10句话

func main() {
	filePath := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\13_file\\05_fileWrite03\\test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件失败，错误为：", err)
		return
	}
	//关闭文件
	defer file.Close()

	str := "覆盖之后的你好\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}
