package process

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/client/utils"
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

// 发送群聊的消息
func (sp *SmsProcess) SendGroupMes(content string) (err error) {

	//创建mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//创建SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(smsMes) err =  ", err)
		return
	}

	mes.Data = string(data)

	//对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(mes) err =  ", err)
		return
	}

	//把mes发给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes tf.WritePkg(data) err = ", err)
		return
	}
	return
}
