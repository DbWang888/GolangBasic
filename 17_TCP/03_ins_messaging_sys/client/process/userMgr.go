package process

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/client/model"
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //在登录成功后 完成对CurUser的初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, user := range onlineUsers {
		fmt.Printf("用户id：%v\t Name:\t%+v\n", id, user)
	}
}

// 处理返回的NotifyUserStatusMes
func updataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//适当优化，原先是否已存在
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
