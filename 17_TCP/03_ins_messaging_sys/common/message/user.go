package message

//用户信息结构体，字段名使用tag与序列化后的json字符串对应
type User struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
}
