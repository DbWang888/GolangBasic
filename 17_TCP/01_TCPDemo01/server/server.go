package main

import (
	"fmt"
	"net"
)

//为客户端服务协程，
func process(conn net.Conn) {

	//关闭conn
	defer conn.Close()

	//循环接收客户端发送的数据
	for {
		buf := make([]byte, 1024) //创建一个切片

		//等待客户端通过conn发送信息
		//如果客户端没有Write，则协程阻塞在这里
		fmt.Println("服务器等待接收信息 " + conn.RemoteAddr().String())
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("服务器read错误 err = ", err)
			return
		}

		//显示客户端发送的内容
		fmt.Print(string(buf[:n]))

	}

}

func main() {
	fmt.Println("服务端开始监听......")
	//网络协议tcp，本地监听9999端口
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("监听失败 err = ", err)
		return //监听失败则无后续，直接退出函数
	}
	defer listener.Close() //最后关闭监听

	//循环等待客户端链接我
	for {
		fmt.Println("等待客户端连接")
		//等待客户端连接，未连接则一直阻塞
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("当前连接失败 err = ", err)
		} else {
			fmt.Printf("连接成功 conn = %v,客户端IP = %v \n", conn, conn.RemoteAddr().String())
		}
		//开启一个协程为客户端服务
		go process(conn)
	}

	// fmt.Printf("listen suc = %v\n", listener)

}
