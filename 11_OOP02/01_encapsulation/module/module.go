package module

import "fmt"

//封装person结构体
type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(person_age int) {
	if person_age > 0 && person_age < 150 {
		p.age = person_age
		fmt.Println("年龄保存成功")
	} else {
		fmt.Println("年龄输入错误")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(money float64) {
	if money > 0 {
		p.salary = money
		fmt.Println("工资保存成功")
	} else {
		fmt.Println("工资输入错误")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}

//封装account结构体

type account struct {
	account_num string
	banlance    float64
	pwd         string
}

//得到一个结构体指针
func SetNewAccount() *account {
	//初始密码为666666
	return &account{
		pwd: "666666",
	}

}

//设置账号
func (a *account) SetAccountNum(acc_num string, pwd string) {
	if pwd == a.pwd {
		if len(acc_num) >= 6 && len(acc_num) <= 10 {
			a.account_num = acc_num
			fmt.Println("账号设置成功")
		} else {
			fmt.Println("账号输入错误")
		}
	} else {
		fmt.Println("账号输入错误")
	}

}

//查看账号
func (a *account) GetAccountNum() string {
	return a.account_num
}

//设置余额
func (a *account) SetBanlance(banl float64, pwd string) {
	if pwd == a.pwd {
		if banl > 20.0 {
			a.banlance = banl
			fmt.Println("余额设置成功")
		} else {
			fmt.Println("余额设置失败")
		}
	} else {
		fmt.Println("账号输入错误")
	}
}

//查看余额
func (a *account) GetBanlance(pwd string) float64 {
	if pwd == a.pwd {
		return a.banlance
	} else {
		fmt.Println("密码错误")
		return 0.0
	}
}

//修改密码

func (a *account) SetPwd(pwd string, newPwd string) {

	if pwd == a.pwd {
		if len(pwd) == 6 && pwd != newPwd {
			a.pwd = newPwd
			fmt.Println("密码设置成功")
		} else {
			fmt.Println("密码设置失败。密码长度为6位且不能与原密码相同")
		}
	}

}
