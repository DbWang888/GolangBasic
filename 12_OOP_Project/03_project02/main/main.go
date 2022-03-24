package main

import (
	"12_OOP_Project/03_project02/customer"
	"12_OOP_Project/03_project02/service"
	"fmt"
)

//界面层

//表示软件界面必要内容
type customerView struct {
	key  string //获取输入选项
	loop bool   //是否退出  false 为退出
	//customerService
	customerService *service.Service
}

func (c *customerView) customerList() {

	//获取当前所有的客户信息
	customers := c.customerService.List()
	fmt.Println("------------------------------客户列表------------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------------------------客户列表完成------------------------------")
}

//添加客户方法，调用Service层的添加方法
func (c *customerView) customerAdd() {
	fmt.Println("------------------------------添加客户------------------------------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	//此处不输入Id，id是唯一的，需要系统分配
	cusNoneId := customer.NewCustomerNoneId(name, gender, age, phone, email)
	if c.customerService.Add(cusNoneId) {
		fmt.Println("------------------------------添加成功------------------------------")
	} else {
		fmt.Println("添加失败")
	}
}

//根据id删除客户
func (c *customerView) customerDelete() {
	fmt.Println("请输入要删除的客户Id")
	id := 0
	fmt.Scanln(&id)
	if c.customerService.Delete(id) {
		fmt.Println("删除成功")
	} else {
		fmt.Println("未找到相应的客户，删除失败")
	}
}

//根据id修改客户信息

func (c *customerView) customerChange() {
	fmt.Println("----------------------------客户信息修改-----------------------------")
	fmt.Println("请输入要修改的客户id")
	id := 0
	fmt.Scanln(&id)
	index := c.customerService.FindById(id)
	if index == -1 {
		fmt.Println("当前id不存在，请重新输入")
		return
	}
	fmt.Println("请输入修改后的姓名，回车表示不修改")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入修改后的性别，回车表示不修改")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入修改后的年龄，回车表示不修改")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入修改后的电话，回车表示不修改")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("请输入修改后的邮箱，回车表示不修改")
	email := ""
	fmt.Scanln(&email)
	c.customerService.Change(index, name, gender, age, phone, email)
}

func (c *customerView) mainView() {
	for {
		//软件文字界面
		fmt.Println("------------------------客户信息管理软件------------------------")
		fmt.Println("             1 添加客户")
		fmt.Println("             2 修改客户")
		fmt.Println("             3 删除客户")
		fmt.Println("             4 客户列表")
		fmt.Println("             5 退出软件")
		fmt.Println("             ")
		fmt.Print("            请选择（1-5）：")
		fmt.Scanln(&(c.key))
		switch c.key {
		case "1":
			c.customerAdd()
		case "2":
			c.customerChange()
		case "3":
			c.customerDelete()
		case "4":
			c.customerList()
		case "5":
			c.loop = true
			fmt.Println("退出软件")
		default:
			fmt.Println("您输入的内容有误，请重新输入")
		}

		if c.loop == true {
			break
		}
	}

}

func main() {
	a := customerView{
		key:             "",
		loop:            false,
		customerService: service.NewCustomerService(),
	}
	a.mainView()
}
