package main

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/client/process"
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id 一个表示系统变量
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	//判断是否还继续显示菜单
	var key int
	// var loop = true

	for {
		fmt.Println("----------欢迎登陆即时通讯聊天系统----------")
		fmt.Println("\t\t 1 登录聊天系统")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择（1-3）：")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的Id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			//完成登录
			//创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			// loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请设置用户的Id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请设置用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请设置用户的昵称")
			fmt.Scanf("%s\n", &userName)

			//创建一个UserProcess的实例，完成注册请求
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)

		case 3:
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0)
		default:
			fmt.Println("输入错误，重新输入")
		}
	}

}
