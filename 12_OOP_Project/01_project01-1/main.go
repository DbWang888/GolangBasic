package main

import "fmt"

func main() {
	//显示主菜单
	//功能对应序号
	var key string
	//是否退出转换
	var loop bool = true
	//账户金额和收支金额
	var accountMoney, recAndExpMoney float64
	//收入支出说明及明细
	var note, details string
	details = "收支\t账户金额\t\t收支金额\t\t说明\t"
	for {
		fmt.Println("-----------------------家庭收支记账软件-----------------------")
		fmt.Println()
		fmt.Println("                      1 收支明细")
		fmt.Println("                      2 登记收入")
		fmt.Println("                      3 登记支出")
		fmt.Println("                      4 退出软件")
		fmt.Println()
		fmt.Print("                      请选择（1-4）：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("\n-----------------------当前收支明细记录-----------------------")
			fmt.Println(details)
		case "2":
			fmt.Print("登记收入金额")
			fmt.Scanln(&recAndExpMoney)
			accountMoney += recAndExpMoney
			fmt.Print("登记收入说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%.2f\t\t%.2f\t\t%v\t", accountMoney, recAndExpMoney, note)
		case "3":
			fmt.Println("登记支出金额")
			fmt.Scanln(&recAndExpMoney)
			if recAndExpMoney > accountMoney {
				fmt.Println("余额不足，支出失败")
			} else {
				accountMoney -= recAndExpMoney
				fmt.Println("登记支出说明")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n支出\t%.2f\t\t%.2f\t\t%v\t", accountMoney, recAndExpMoney, note)
			}
		case "4":
			var confirm string
			fmt.Print("请确认是否退出软件？ y/n")
			for {
				fmt.Scanln(&confirm)
				if confirm == "y" {
					loop = false
					break
				} else if confirm == "n" {
					break
				} else {
					fmt.Println("输入错误请重新输入")
				}
			}

		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
		if loop == false {
			fmt.Println("退出软件")
			break
		}

	}
}
