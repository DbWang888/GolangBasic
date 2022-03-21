package main

import (
	"11_OOP02/01_encapsulation/module"
	"fmt"
)

func main() {
	// var person1 module.Person = module.Person{
	// 	Name: "小明",
	// }

	person1 := module.NewPerson("小明")
	(*person1).SetAge(20)
	(*person1).SetSalary(5663.69)
	age := person1.GetAge()
	sal := person1.GetSalary()

	fmt.Println(*person1)
	fmt.Println(age)
	fmt.Println(sal)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	//新账号
	newAcc := module.SetNewAccount()
	fmt.Println(newAcc)
	newAcc.SetBanlance(665.36, "666666")
	newAcc.SetPwd("666666", "123456")
	newAcc.SetAccountNum("666666", "123456")
	fmt.Printf("%+v\n", *newAcc)
}
