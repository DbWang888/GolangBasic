package customerView

import (
	"12_OOP_Project/04_project02-2/customerService"
	"fmt"
)

//软件界面结构体
type customerView struct {
	key     string                           // 接收输入的key值，匹配功能选项
	loop    bool                             //是否退出软件参数
	Service *customerService.CustomerService //取customerservice结构体地址 方便调用方法
}

//删除客户信息
func (c *customerView) deleteCustomer() {
	fmt.Println("--------------------删除客户信息--------------------")
	fmt.Print("请输入删除客户的id(-2退出)：")
	id := 0
	fmt.Scanln(&id)
	if id == -2 {
		return
	}
	index := c.Service.FindById(id)
	if index == -1 {
		fmt.Println("输入错误，无当前客户")
		return
	}
	for {
		fmt.Print("确认是否删除（y/n）：")
		key := ""
		fmt.Scanln(&key)
		if key == "y" {
			c.Service.Delete(index)
			return
		} else if key == "n" {
			return
		} else {
			fmt.Println("输入错误，重新输入")
		}
	}

}

//对客户信息进行修改
func (c *customerView) ChangeCustomerInfo() {
	fmt.Println("--------------------修改客户信息--------------------")
	for {
		fmt.Print("请输入待修改客户id（-2退出）：")
		id := 0
		fmt.Scanln(&id)
		if id == -2 {
			return
		}
		index := c.Service.FindById(id)
		if index == -1 {
			fmt.Println("输入Id错误，无法找到对应客户，请重新输入")
		} else {
			fmt.Print("请输入客户姓名(直接回车表示不修改)：")
			name := ""
			fmt.Scanln(&name)
			fmt.Print("请输入客户性别：")
			gender := ""
			fmt.Scanln(&gender)
			fmt.Print("请输入客户年龄：")
			age := 0
			fmt.Scanln(&age)
			fmt.Print("请输入客户电话：")
			phone := ""
			fmt.Scanln(&phone)
			fmt.Print("请输入客户邮箱：")
			email := ""
			fmt.Scanln(&email)
			c.Service.ChangeInfo(index, name, gender, age, phone, email)
			break
		}
	}

}

//遍历输出客户信息，显示所有客户
func (c *customerView) ListCustomer() {
	fmt.Println("--------------------客户信息列表--------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	fmt.Println(c.Service.Customers)
	for i := 0; i < len(c.Service.Customers); i++ {
		// if c.Service.Customers[i].Id == 0 {
		// 	return
		// }
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t\t%v\n", c.Service.Customers[i].Id, c.Service.Customers[i].Name,
			c.Service.Customers[i].Gender, c.Service.Customers[i].Age, c.Service.Customers[i].Phone, c.Service.Customers[i].Email)
	}
}

//添加客户的界面
func (c *customerView) AddCustomer() {
	fmt.Println("--------------------添加客户信息--------------------")
	fmt.Print("请输入客户姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("请输入客户性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("请输入客户年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("请输入客户电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("请输入客户邮箱：")
	email := ""
	fmt.Scanln(&email)
	c.Service.Add(name, gender, age, phone, email)
	fmt.Println("--------------------添加成功--------------------")
}

//项目主界面
func (c *customerView) MainView() {
	for {
		fmt.Println("--------------------客户信息管理软件--------------------")
		fmt.Println("                     1 添加客户")
		fmt.Println("                     2 修改客户")
		fmt.Println("                     3 删除客户")
		fmt.Println("                     4 客户列表")
		fmt.Println("                     5 退出软件")
		fmt.Println()
		fmt.Print("                     请选择(1-5): ")
		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			c.AddCustomer()
		case "2":
			c.ChangeCustomerInfo()
		case "3":
			c.deleteCustomer()
		case "4":
			c.ListCustomer()
		case "5":
			for {
				key := ""
				fmt.Println("请确认是否退出？ y/n")
				fmt.Scanln(&key)
				if key == "y" {
					c.loop = true
					break
				} else if key == "n" {
					break
				} else {
					fmt.Println("输入错误，重新输入")
				}
			}
		default:
			fmt.Println("输入错误，请重新输入")
		}

		//判断参数，是否退出软件
		if c.loop == true {
			fmt.Println("------------------客户信息管理软件退出------------------")
			break
		}
	}
}

func NewMainView() *customerView {
	return &customerView{
		key:     "",
		loop:    false,
		Service: &customerService.CustomerService{},
	}
}
