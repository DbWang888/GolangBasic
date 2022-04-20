package utils

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//这里将这些方法关联到结构体中
//传输相关的内容
type Transfer struct {
	Conn net.Conn   //连接
	Buf  [8096]byte //传输时使用的缓冲
}

//处理客户端的信息包
func (r *Transfer) ReadPkg() (mes message.Message, err error) {

	// buf := make([]byte, 8096)
	fmt.Println("等待读取服务器发送的数据")
	//在conn未关闭的情况下才会阻塞，如果客户端关闭了conn则不会阻塞
	_, err = r.Conn.Read(r.Buf[0:4]) //读不到东西会阻塞
	if err != nil {
		fmt.Println("client func readpak conn.Read err = ", err)
		// err = errors.New("read pkg header error")
		return
	}
	//根据buf[:4] 转换成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(r.Buf[0:4])

	//根据pkgLen 读取消息内容
	n, err := r.Conn.Read(r.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("r.Conn.Read(r.Buf[:4]) err = ", err)
		// err = errors.New("read pkg body error")
		return
	}

	//把pkgLen反序列化成 message.Message类型
	err = json.Unmarshal(r.Buf[:pkgLen], &mes)

	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) err = ", err)
		return
	}
	return
}

//发送消息函数
func (r *Transfer) WritePkg(data []byte) (err error) {
	//发送一个长度给对方
	fmt.Println("conn wrtie = ", r.Conn.LocalAddr().String())
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(r.Buf[0:4], pkgLen)
	n, err := r.Conn.Write(r.Buf[:4])
	if err != nil || n != 4 {
		fmt.Println("r.Conn.Write(bytes) err = ", err)
		return
	}
	fmt.Printf("消息长度发送成功，长度为：%v 内容为 %s\n", pkgLen, data)

	//发送消息本身
	n, err = r.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("r.Conn.Write(data) err = ", err)
		return
	}

	return
}
