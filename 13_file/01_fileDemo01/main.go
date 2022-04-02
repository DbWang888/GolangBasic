package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//带缓冲区的方式读取文件内容

func main() {

	//打开文件
	//概念说明：
	//1.file 叫 file对象、file指针、file文件句柄
	file, err := os.Open("C:/Users/wangli/Desktop/test/name.txt")
	if err != nil {
		fmt.Println("打卡文件错误：", err)
	}
	//输出文件，可看出file就是指针
	fmt.Printf("file = %v\n", file)

	//关闭文件  defer函数，在函数退出时关闭文件，不影响后续操作
	defer file.Close()

	//创建一个*Reader 带缓冲的

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束，意思是按行读取
		if err == io.EOF {                  //io.EOF  表示文件的末尾
			break
		}
		//输出内容
		fmt.Println(str)

	}

}
