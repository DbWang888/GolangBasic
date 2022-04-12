//客户端
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	defer conn.Close() //客户端也需要conn.Close()
	if err != nil {
		fmt.Println("连接失败 err = ", err)
		return
	}
	fmt.Println("连接成功 conn = ", conn)
	//客户端可发送单行数据后退出
	reader := bufio.NewReader(os.Stdin)

	for {

		//从终端读取一行用户输入，并准备发给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader读取失败， err = ", err)
		}
		line = strings.Trim(line, " \r\n")

		if line == "exit" {
			return
		}
		//再将line发送给服务器
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("line发送给服务器失败 err = ", err)
		}
		fmt.Printf("客户端发送了%d个数据\n", n)
	}
}
