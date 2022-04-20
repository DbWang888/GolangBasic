package process

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/client/utils"
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

//转发消息
func (sp *SmsProcess) SendGroupMes(mes *message.Message) {
	
	//先取出mes中的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal([]byte(mes.Data),&smsMes) err = ",err)
		return
	}

	data, err := json.Marshal(mes)

	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(mes) err = ",err)
		return
	}

	//遍历服务器端的onlineUser 将消息转发出去
	for id, up := range UserMgr.onlineUsers {

		//需要过滤掉自己，不发给自己
		if id == smsMes.UserId {
			continue
		}

		sp.sendMesToEachOnlineUser(data,up.Conn)
	}




}

//转发给每一个在线的用户
func (sp *SmsProcess) sendMesToEachOnlineUser(data []byte, 
	conn net.Conn){
	//创建tf实例 发送data
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)

	if err != nil {
		fmt.Println("sendMesToEachOnlineUser tf.WritePkg(data) err = ",err)
		return
	}

}