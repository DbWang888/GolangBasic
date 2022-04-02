package main

import (
	"bufio"
	"fmt"
	"os"
)

//带缓存的创建文件并写入内容

func main() {

	filePath := "C:/Users/wangli/Desktop/test/testWrite.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("openfile err = %v", err)
		return
	}

	defer file.Close()

	str := "你好\n"

	// 写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer带缓存，调用writerString方法时，是先将内容写入缓存
	//需要调用Flush方法将缓存的数据写入文件中，否则文件内没有数据、
	writer.Flush()
}
