package process

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"encoding/json"
	"fmt"
)

//显示输出信息
func outputGroupMes(mes *message.Message) {

	//反序列化mes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes json.Unmarshal([]byte(mes.Data),&smsMes) err = ", err)
		return
	}

	//显示信息格式化
	info := fmt.Sprintf("用户id：%d 说：\t%s\n", smsMes.UserId, smsMes.Content)

	fmt.Println(info)
	fmt.Println()

}
