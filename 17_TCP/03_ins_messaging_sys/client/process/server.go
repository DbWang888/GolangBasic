package process

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/client/utils"
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面
func ShowMenu() {
	fmt.Println("---------恭喜XXX登录成功----------")
	fmt.Println("\t\t1 显示在线用户列表")
	fmt.Println("\t\t2 发送消息")
	fmt.Println("\t\t3 信息列表")
	fmt.Println("\t\t4 退出系统")
	fmt.Println("\t\t请选择（1-4）")
	var key int
	var content string
	fmt.Scanf("%d\n", &key)
	//总会发送消息，将其定义在switch外，避免重复
	sm := SmsProcess{}
	switch key {
	case 1:
		// fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息")
		fmt.Println("请输入你要发送的内容：")
		fmt.Scanf("%s\n", &content)

		sm.SendGroupMes(content)

	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误，重新输入")
	}

}

//和服务器端保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("client func serverProcessMes read err = ", err)
			return
		}

		//如果读取到消息
		//fmt.Println("serverProcessMes mes = ", mes)
		switch mes.Type {

		case message.NotifyUserStatusMesType:

			fmt.Println("NotifyUserStatusMesType = ", mes.Type)
			//取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			fmt.Println(" json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes) = ", err)
			//把用户信息保存到客户map[int]User中
			updataUserStatus(&notifyUserStatusMes)

		case message.SmsMesType: //群发消息
			outputGroupMes(&mes)

		default:
			fmt.Println("服务端返回未知消息，该类型暂时无法处理")
		}
	}
}
