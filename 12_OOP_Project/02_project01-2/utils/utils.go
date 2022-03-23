package utils

import (
	"fmt"
)

//记账软件需要的字段结构体
type account struct {
	key            string //获取输入值
	loop           bool   //是否退出决定值
	details        string //接收记账详情
	haveDetails1   string
	haveDetails2   string
	note           string  //收入支出说明
	recAndExpMoney float64 //收入支出金额
	accountMoney2  float64 //总金额
	accountMoney1  float64 //总金额
	currentAccount string  //当前账户
	changeMoney    float64 //转账金额
}

//转账
func (a *account) change() {
	fmt.Printf("\n-----------------------当前账户为%v-----------------------", a.currentAccount)
	if a.currentAccount == "账户1" {
		fmt.Print("输入转出金额")
		fmt.Scanln(&a.changeMoney)
		a.accountMoney1 -= a.changeMoney
		a.accountMoney2 += a.changeMoney
		a.haveDetails1 += fmt.Sprintf("\n支出\t%.2f\t\t%.2f\t\t转出\t", a.accountMoney1, a.changeMoney)
		a.haveDetails2 += fmt.Sprintf("\n收入\t%.2f\t\t%.2f\t\t转入\t", a.accountMoney2, a.changeMoney)
	} else {
		fmt.Print("输入转出金额")
		fmt.Scanln(&a.changeMoney)
		a.accountMoney2 -= a.changeMoney
		a.accountMoney1 += a.changeMoney
		a.haveDetails2 += fmt.Sprintf("\n支出\t%.2f\t\t%.2f\t\t转出\t", a.accountMoney1, a.changeMoney)
		a.haveDetails1 += fmt.Sprintf("\n收入\t%.2f\t\t%.2f\t\t转入\t", a.accountMoney2, a.changeMoney)
	}
}

//切换账户
func (a *account) changeAccount() {
	fmt.Printf("\n-----------------------当前账户为%v-----------------------", a.currentAccount)
	if a.currentAccount == "账户1" {
		a.currentAccount = "账户2"
	} else {
		a.currentAccount = "账户1"
	}
	fmt.Printf("账户切换成功，当前账户为%v", a.currentAccount)
}

//返回结构体地址，且初始化字段
func GetAccount() *account {
	return &account{
		key:            "",
		loop:           true,
		details:        "\n收支\t账户金额\t\t收支金额\t\t说明\t",
		haveDetails1:   "",
		haveDetails2:   "",
		note:           "",
		recAndExpMoney: 0.0,
		accountMoney1:  10000.0,
		accountMoney2:  10000.0,
		currentAccount: "账户1",
	}
}

//显示收支明细记录方法
func (a *account) showDetails() {
	fmt.Printf("\n-----------------------%v收支明细记录-----------------------", a.currentAccount)
	if a.currentAccount == "账户1" {
		if a.haveDetails1 == "" {
			fmt.Println("当前暂无记录")
		} else {
			fmt.Println(a.details + a.haveDetails1)
		}
	} else {
		if a.haveDetails2 == "" {
			fmt.Println("当前暂无记录")
		} else {
			fmt.Println(a.details + a.haveDetails2)
		}
	}
}

//记录收入方法
func (a *account) getIncome() {
	if a.currentAccount == "账户1" {
		fmt.Print("登记收入金额")
		fmt.Scanln(&a.recAndExpMoney)
		a.accountMoney1 += a.recAndExpMoney
		fmt.Print("登记收入说明")
		fmt.Scanln(&a.note)
		a.haveDetails1 += fmt.Sprintf("\n收入\t%.2f\t\t%.2f\t\t%v\t", a.accountMoney1, a.recAndExpMoney, a.note)
	} else {
		fmt.Print("登记收入金额")
		fmt.Scanln(&a.recAndExpMoney)
		a.accountMoney2 += a.recAndExpMoney
		fmt.Print("登记收入说明")
		fmt.Scanln(&a.note)
		a.haveDetails2 += fmt.Sprintf("\n收入\t%.2f\t\t%.2f\t\t%v\t", a.accountMoney2, a.recAndExpMoney, a.note)
	}
}

//记录支出方法
func (a *account) expenditure() {
	if a.currentAccount == "账户1" {
		fmt.Println("登记支出金额")
		fmt.Scanln(&a.recAndExpMoney)
		if a.recAndExpMoney > a.accountMoney1 {
			fmt.Println("余额不足，支出失败")
		} else {
			a.accountMoney1 -= a.recAndExpMoney
			fmt.Println("登记支出说明")
			fmt.Scanln(&a.note)
			a.haveDetails1 += fmt.Sprintf("\n支出\t%.2f\t\t%.2f\t\t%v\t", a.accountMoney1, a.recAndExpMoney, a.note)
		}
	} else {
		fmt.Println("登记支出金额")
		fmt.Scanln(&a.recAndExpMoney)
		if a.recAndExpMoney > a.accountMoney2 {
			fmt.Println("余额不足，支出失败")
		} else {
			a.accountMoney2 -= a.recAndExpMoney
			fmt.Println("登记支出说明")
			fmt.Scanln(&a.note)
			a.haveDetails2 += fmt.Sprintf("\n支出\t%.2f\t\t%.2f\t\t%v\t", a.accountMoney2, a.recAndExpMoney, a.note)
		}
	}

}

//退出方法
func (a *account) quit() {
	var confirm string
	fmt.Print("请确认是否退出软件？ y/n")
	for {
		fmt.Scanln(&confirm)
		if confirm == "y" {
			a.loop = false
			break
		} else if confirm == "n" {
			break
		} else {
			fmt.Println("输入错误请重新输入")
		}
	}
}

//主界面方法
func (a *account) GetMainInterface() {
	for {
		fmt.Println("-----------------------家庭收支记账软件-----------------------")
		fmt.Println()
		fmt.Println("                      1 收支明细")
		fmt.Println("                      2 登记收入")
		fmt.Println("                      3 登记支出")
		fmt.Println("                      4 转账")
		fmt.Println("                      5 退出软件")
		fmt.Println("                      6 切换账户")
		fmt.Println()
		fmt.Print("                      请选择（1-6）：")
		fmt.Scanln(&a.key)
		switch a.key {
		case "1":
			a.showDetails()
		case "2":
			a.getIncome()
		case "3":
			a.expenditure()
		case "4":
			a.change()
		case "5":
			a.quit()
		case "6":
			a.changeAccount()
		}
		fmt.Println()
		if a.loop == false {
			fmt.Println("退出软件")
			break
		}

	}
}
