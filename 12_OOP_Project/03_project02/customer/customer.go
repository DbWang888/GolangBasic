package customer

import "fmt"

//数据层

//声明一个customer结构体，表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//工厂模式返回customer实例指针
func NewCustomer(id int, name string, gender string, age int,
	phone string, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//工厂模式返回customer实例指针,不带ID
func NewCustomerNoneId(name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//工厂模式 返回该客户信息的的格式化字符串，
//如“1       张三    男      23              3302252         zhangsan@qq.com
func (c Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t\t%v\t\t%v", c.Id, c.Name,
		c.Gender, c.Age, c.Phone, c.Email)
	return info
}
