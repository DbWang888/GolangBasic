package main

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"03_ins_messaging_sys/03_ins_messaging_sys/server/process"
	"03_ins_messaging_sys/03_ins_messaging_sys/server/utils"
	"fmt"
	"io"
	"net"
)

//先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//根据客户端发送的消息种类不同，决定调用哪个函数来处理
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	//看看是否接受到客户端群发的消息
	// fmt.Println("MES = ", mes)
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑函数
		// 创建一个UserProcess实例
		up := process.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		// 创建一个UserProcess实例
		up := &process.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//处理群发消息
		sp := &process.SmsProcess{}

		sp.SendGroupMes(mes)



	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (p *Processor) process2() (err error) {
	//循环读客户端发出的消息
	//创建一个transfer实例完成读包任务
	tf := &utils.Transfer{
		Conn: p.Conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端退出")
				return err
			} else {
				fmt.Println("func process2 readPkg(conn) err = ", err)
				return err
			}
		}
		fmt.Println("mes process2 = ", mes)
		err = p.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("mes process2 ERR = ", err)
			return err
		}
	}
}
