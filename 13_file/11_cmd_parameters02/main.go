package main

import (
	"flag"
	"fmt"
)

func main() {
	//定义几个变量，接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 88, "端口号，默认为88")
	flag.Parse()

	fmt.Println(user, pwd, host, port)

}
