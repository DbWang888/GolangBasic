package message

//确定消息常量类型
const (
	LoginMesType            = "LoginMes"            //登录发送息的消息
	LoginResMesType         = "LoginResMes"         //登录，服务端返回消息
	RegisterMesType         = "RegisterMes"         //注册发送息的消息
	RegisterResMesType      = "RegisterResMes"      //注册，服务端返回消息
	NotifyUserStatusMesType = "NotifyUserStatusMes" //用户状态
	SmsMesType              = "SmsMes"              //发送消息
)

//定义用户状态的常量
const (
	UserOnline = iota //下面会自动增长
	UserOffline
)

//消息的结构体
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的信息
}

//登录发送息的消息
type LoginMes struct {
	UserId   int    `json:"useriId"`  //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

//登录，服务端返回消息
type LoginResMes struct {
	Code    int    `json:"code"`    //返回状态码  500表示用户未注册    200表示登录成功
	UsersId []int  `json:"usersid"` //增加字段,保存用户id切片
	Error   string `json:"error"`   //返回错误信息
}

//
type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码  400该用户已占用    200表示注册成功
	Error string `json:"error"` //返回错误信息
}

//配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户的状态

}

//增加smsMes  发送的消息
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体 继承
}

