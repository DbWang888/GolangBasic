package process

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"03_ins_messaging_sys/03_ins_messaging_sys/server/model"
	"03_ins_messaging_sys/03_ins_messaging_sys/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int //连接对应的id
}

//通知所有在线的用户某一个用户上线
//userid要通知别的在线用户，自己上线
func (u *UserProcess) NotifyOthersOnlineUser(userId int) {

	//遍历OnlineUsers，然后一个一个发送NotifyUserStatusMes
	for id, up := range UserMgr.onlineUsers {
		if id == userId {
			continue
		}

		//开始通知（单独写方法）
		up.NotifyMeOnline(userId)
	}

}

//通知
func (u *UserProcess) NotifyMeOnline(userId int) {

	//组装消息NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes进行json序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Masrshal(notifyUserStatusMes) err = ", err)
		return
	}

	//将序列化后的notifyUserStatusMes 赋值给mes.Data
	mes.Data = string(data)
	//对Mes进行再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Masrshal(mes) err = ", err)
		return
	}

	//发送
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(NotifyMeOnline) err = ", err)
		return
	}

}

//处理注册请求
func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	//从mes中取出mes.Data，反序列化成registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), registerMes) err = ", err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//再声明一个loginResMes，并完成赋值
	var registerResMes message.RegisterResMes

	//需要到redis数据库去完成注册
	//使用model.MyUserDao 到redis中去完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	fmt.Println("registerResMes.code= ", registerResMes.Code)

	//对resMes进行序列化，准备发送
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) err = ", err)
		return
	}
	//将data值赋给resMes
	resMes.Data = string(data)

	//对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err = ", err)
		return
	}

	//发送data，将代码封装到writePkg函数中
	//创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	fmt.Println("ServerProcessRegister 发送成功")
	if err != nil {
		fmt.Println("ServerProcessRegister tf.WritePkg(data) err =  ", err)
	}
	return

}

//编写一个函数，专门处理登录请求
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	//登录核心代码
	//从mes中取出mes.Data，反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), loginMes) err = ", err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//再声明一个loginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//需要到redis数据库去完成验证
	//使用model.MyUserDao 到redis中去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200

		//用户登录成功，把该登录成功的用户放到UserMgr中去
		//将登录成功的用户的id赋值给u.UserId
		u.UserId = loginMes.UserId
		UserMgr.AddOnlineUser(u)
		// 通知用户我上线了
		u.NotifyOthersOnlineUser(loginMes.UserId)

		//遍历userMgr.onlineUsers取出id
		for id := range UserMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Printf("%s登录成功\n", user.UserName)
	}

	//将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) err = ", err)
		return
	}
	//将data赋值给resMes
	resMes.Data = string(data)
	fmt.Println("mes data = ", string(data))
	//对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err = ", err)
		return
	}

	//发送data，将代码封装到writePkg函数中
	//创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)

	return
}
