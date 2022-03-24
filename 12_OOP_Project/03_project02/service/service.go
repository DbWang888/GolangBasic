package service

import (
	"12_OOP_Project/03_project02/customer"
)

//业务逻辑层

//完成对customer的操作，增删改查
//通过结构体切片完成
type Service struct {
	Customers   []customer.Customer //customer结构体切片
	CustomerNum int                 //表示当前切片客户数量，后期表示+1为id
}

//返回一个service实例
func NewCustomerService() *Service {
	customerService := &Service{}
	customerService.CustomerNum = 1
	customerStruct := customer.NewCustomer(1, "张三", "男", 23, "3302252", "zhangsan@qq.com")
	customerService.Customers = append(customerService.Customers, *customerStruct)
	return customerService
}

//返回客户信息切片
func (s *Service) List() []customer.Customer {
	// var list string = ""
	// for i := 0; i < len(s.Customers); i++ {
	// 	list += fmt.Sprintf("%v\t%v\t%v\t%v\t\t%v\t\t%v", s.Customers[i].Id, s.Customers[i].Name,
	// 		s.Customers[i].Gender, s.Customers[i].Age, s.Customers[i].Phone, s.Customers[i].Email)
	// }
	return s.Customers
}

//新增客户信息(添加切片 )
func (s *Service) Add(customerid customer.Customer) bool {
	s.CustomerNum++
	customerid.Id += s.CustomerNum
	s.Customers = append(s.Customers, customerid)
	return true
}

//删除客户方法
func (s *Service) Delete(id int) bool {
	index := s.FindById(id)
	if index != -1 {
		s.Customers = append(s.Customers[:index], s.Customers[index+1:]...)
		return true
	}
	return false
}

//根据id找对应下标方法
func (s *Service) FindById(id int) int {
	index := -1
	for i := 0; i < len(s.Customers); i++ {
		if s.Customers[i].Id == id {
			index = i
			return index
		}
	}
	return index
}

//修改客户方法
func (s *Service) Change(index int, name string, gender string, age int,
	phone string, email string) int {
	i := 0
	if name != "" {
		s.Customers[index].Name = name
		i++
	}
	if gender != "" {
		s.Customers[index].Gender = gender
		i++
	}
	if age != 0 {
		s.Customers[index].Age = age
		i++
	}
	if phone != "" {
		s.Customers[index].Phone = phone
		i++
	}
	if email != "" {
		s.Customers[index].Email = email
		i++
	}
	return i
}
