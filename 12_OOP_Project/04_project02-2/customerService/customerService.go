package customerService

//业务逻辑层  提供对客户信息的增删改查操作

import (
	"12_OOP_Project/04_project02-2/model"
)

// 业务逻辑结构体
type CustomerService struct {
	Customers  []model.Customer //客户信息结构体切片，用于存储客户信息
	CustomerId int
}

//删除客户
func (c *CustomerService) Delete(index int) {
	// if len(c.Customers) == 1 {
	// 	c.Customers = make([]model.Customer, 1)
	// 	return
	// }
	// if index == (len(c.Customers) - 1) {
	// 	c.Customers = c.Customers[:index]
	// 	return
	// }
	c.Customers = append(c.Customers[:index], c.Customers[index+1:]...)
}

//通过id找对应的切片index
func (c *CustomerService) FindById(id int) int {
	for i := 0; i < len(c.Customers); i++ {
		if c.Customers[i].Id == id {
			return i
		}
	}
	return -1
}

//对客户信息进行修改
func (c *CustomerService) ChangeInfo(index int, name string, gender string, age int, phone string,
	email string) {
	changinfo := ""
	if name != "" {
		c.Customers[index].Name = name
		changinfo += "姓名"
	}
	if gender != "" {
		c.Customers[index].Gender = gender
		changinfo += "性别"
	}
	if age != 0 {
		c.Customers[index].Age = age
		changinfo += "年龄"
	}
	if phone != "" {
		c.Customers[index].Phone = phone
		changinfo += "电话"
	}
	if email != "" {
		c.Customers[index].Email = email
		changinfo += "邮箱"
	}
}

//生成id的方法
func (c *CustomerService) GetId() int {
	if c.CustomerId == 0 {
		c.CustomerId = 1
	} else {
		c.CustomerId++
	}
	return c.CustomerId
}

//调用model的新增客户方法，为客户信息结构体切片新增客户信息
func (c *CustomerService) Add(name string, gender string, age int, phone string,
	email string) {
	customer := model.NewCustomer()
	customer.Name = name
	customer.Id = c.GetId()
	customer.Gender = gender
	customer.Age = age
	customer.Phone = phone
	customer.Email = email
	c.Customers = append(c.Customers, *customer)
}
