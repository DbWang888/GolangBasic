package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//保存统计结果
type Charcount struct {
	ChCount    int //字母个数
	NumCount   int //数字个数
	SpaceCount int //空格个数
	OtherCoutn int //其他字符个数

}

func main() {
	filePath := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\13_file\\09_file_statistics\\test01.txt"

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	defer file.Close()
	charCount := Charcount{}
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v > 'a' && v < 'z':
				fallthrough
			case v > 'A' && v < 'Z':
				charCount.ChCount++
			case v == ' ' || v == '\t':
				charCount.SpaceCount++
			case v >= '0' && v <= '9':
				charCount.NumCount++
			default:
				charCount.OtherCoutn++
			}

		}
	}
	fmt.Printf("%+v", charCount)
}
