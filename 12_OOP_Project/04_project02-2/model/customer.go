package model

//客户信息结构体 含Id、姓名等信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer() *Customer {
	return &Customer{
		Id:     1,
		Name:   "小明",
		Gender: "男",
		Age:    20,
		Phone:  "15220525252",
		Email:  "4242@123.com",
	}
}
