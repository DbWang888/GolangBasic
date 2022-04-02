package testdemo02

import "fmt"

//声明Monster结构体实例，调用store方法
func UseStore() {

	mons := Monster{
		Name:  "理想",
		Age:   22,
		Skill: "吃饭",
	}

	mons.Store()

}

func UseRestore() {
	var monsRe Monster

	monsRe.ReStore()
	fmt.Printf("%+v", monsRe)
}
