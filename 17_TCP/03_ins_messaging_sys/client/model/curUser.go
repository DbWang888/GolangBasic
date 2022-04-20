package model

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"net"
)

//客户端很多地方会使用到curUser，将其作为全局的
type CurUser struct {
	Conn net.Conn
	message.User
}
