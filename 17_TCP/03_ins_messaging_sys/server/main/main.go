package main

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/server/model"
	"fmt"
	"net"
	"time"
)

//处理和客户端的通信
func process1(conn net.Conn) {

	defer conn.Close()

	//这里调用总控，创建
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("go 客户端和服务端通讯协程错误 err= ", err)
		return
	}
}

//编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	//pool为redis.go中的全局变量
	//注意初始化顺序 先initPool 后 initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	//当服务器启动时，先初始化reids连接池
	initPool("127.0.0.1:6379", 16, 0, time.Second*300)
	initUserDao()
}

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	for {
		fmt.Println("等待客户端连接.......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
		}

		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process1(conn)

	}
}
