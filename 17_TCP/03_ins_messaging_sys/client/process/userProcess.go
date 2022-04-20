package process

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/client/utils"
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (u *UserProcess) Register(userId int, userPwd string,
	userName string) (err error) {
	//连接到服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	defer conn.Close()

	//通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//创建registerMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//将registerMes进行json序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("loginMes json.Marshal err = ", err)
		return
	}

	//把data赋给mes.Data字段
	mes.Data = string(data)

	//将mes 进行序列化
	data, err = json.Marshal(mes)

	if err != nil {
		fmt.Println("mes json.Marshal err = ", err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器

	err = tf.WritePkg(data)
	// _, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Println("注册发送信息错误 err = ", err)
		return
	}

	mes, err = tf.ReadPkg() // 就是 RegisResMes

	if err != nil {
		fmt.Println("client  func Register readPkg(conn) err = ", err)
		return
	}

	//将mes的Data部分反序列化为registerResMes
	var registerResMes message.RegisterResMes

	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if err != nil {
		fmt.Println("registerResMes json.Unmarshal([]byte(mes.Data), &registerResMes) err = ", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，重新登录")
		os.Exit(0)

	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}

//写一个函数，完成登录
func (u *UserProcess) Login(userId int, userPwd string) (err error) {
	//下一步开始定协议
	// fmt.Printf("userId = %v userPwd = %v\n", userId, userPwd)
	// return nil

	//连接到服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	// defer conn.Close()

	//通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//创建LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//将loginMes进行json序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes json.Marshal err = ", err)
		return
	}
	//把data赋给mes.Data字段
	mes.Data = string(data)

	//将mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal err = ", err)
		return
	}

	//data就是要发送的json字符串
	//先把data长度发送给服务器
	//获取到data的长度，转换成表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write(bytes) err = ", err)
		return
	}
	fmt.Printf("消息长度发送成功，长度为：%v 内容为 %s\n", pkgLen, string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}

	//休眠20s
	// time.Sleep(time.Second * 20)
	// fmt.Println("休眠了20s")

	//处理服务器端返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	fmt.Printf("abc%+v\n", mes)
	if err != nil {
		fmt.Println("client func Login readPkg(conn) err = ", err)
		return
	}
	//将mes的Data部分反序列化为LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println(" json.Unmarshal([]byte(mes.Data), &loginResMes) err = ", err)
		return
	}
	if loginResMes.Code == 200 {

		// 初始化curUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//可以显示一下当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线列表：")
		for _, v := range loginResMes.UsersId {

			//如果是自己就不显示
			if v == userId {
				continue
			}

			fmt.Println("用户id：\t", v)
			//完成客户端的onlineUsers， 完成初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println()

		//这里还需要一个协程
		//保持和服务器端的通讯，及时接收服务器推送的数据
		go serverProcessMes(conn)

		//循环显示登录成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println("11111111")
		fmt.Println(loginResMes.Error)
	}
	return
}
